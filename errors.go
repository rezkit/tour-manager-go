package tourmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// maxErrorBodyBytes caps how much of a non-2xx response body is buffered
// for error reporting.
const maxErrorBodyBytes = 64 * 1024

// APIError is implemented by every error returned for a non-2xx response.
// Test for it, or for one of the concrete types below, with [errors.As]:
//
//	var verr *tourmanager.ValidationError
//	if errors.As(err, &verr) {
//		// verr.Errors maps field name to validation messages
//	}
//
// APIError is a sealed interface: only types in this package implement it,
// so methods may be added to it in future without that being a breaking
// change for callers.
type APIError interface {
	error

	// StatusCode returns the HTTP status code of the response.
	StatusCode() int

	apiError()
}

// baseError holds the fields common to every APIError. Concrete error
// types embed it by value, which promotes its exported StatusCode method.
type baseError struct {
	Status  int
	Message string
	Method  string
	URL     string
}

func (e *baseError) StatusCode() int { return e.Status }
func (e *baseError) apiError()       {}
func (e *baseError) Error() string {
	return fmt.Sprintf("tourmanager: %s %s: %d %s", e.Method, e.URL, e.Status, e.Message)
}

// ValidationError is returned for a 422 response, when a request failed
// field-level validation.
type ValidationError struct {
	baseError

	// Errors maps a field name to the validation messages for that field.
	Errors map[string][]string
}

// NotFoundError is returned for a 404 response.
type NotFoundError struct {
	baseError
}

// GenericError is returned for a 400, 409 or 429 response, all of which
// the API documents with the same {message: string} shape.
type GenericError struct {
	baseError
}

// UnauthorizedError is returned for a 401 response, meaning the API key
// was missing, invalid, or lacks the scopes required for the request.
type UnauthorizedError struct {
	baseError
}

// UnexpectedStatusError is returned for an undocumented status code, or a
// documented one whose response body didn't match its expected shape.
// Body holds the raw (size-capped) response body for debugging.
type UnexpectedStatusError struct {
	baseError
	Body []byte
}

// RequestError is returned when a request never received a response, for
// example due to a DNS failure, a connection error, or context
// cancellation/timeout.
type RequestError struct {
	Method string
	URL    string
	Err    error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("tourmanager: %s %s: %v", e.Method, e.URL, e.Err)
}
func (e *RequestError) Unwrap() error { return e.Err }

// DecodeError is returned when a 2xx response body didn't decode as the
// shape the client expected.
type DecodeError struct {
	Method string
	URL    string
	Err    error
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("tourmanager: %s %s: decode response: %v", e.Method, e.URL, e.Err)
}
func (e *DecodeError) Unwrap() error { return e.Err }

// newAPIError classifies a non-2xx *http.Response into an APIError. The
// four documented error shapes are treated as applying by status code
// universally, not only on operations that explicitly declare that status
// in openapi.yml (most operations don't bother documenting every status
// they can actually return).
func newAPIError(method, url string, resp *http.Response) error {
	data, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBodyBytes))
	base := baseError{Status: resp.StatusCode, Method: method, URL: url}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		// The spec gives no schema for 401, only an example whose key is
		// "error" (not "message" like every other documented error shape).
		var body struct {
			Error string `json:"error"`
		}
		if json.Unmarshal(data, &body) == nil && body.Error != "" {
			base.Message = body.Error
			return &UnauthorizedError{baseError: base}
		}

	case http.StatusUnprocessableEntity:
		var body struct {
			Message string              `json:"message"`
			Errors  map[string][]string `json:"errors"`
		}
		if json.Unmarshal(data, &body) == nil && body.Message != "" {
			base.Message = body.Message
			return &ValidationError{baseError: base, Errors: body.Errors}
		}

	case http.StatusNotFound:
		var body struct {
			Message string `json:"message"`
		}
		if json.Unmarshal(data, &body) == nil && body.Message != "" {
			base.Message = body.Message
			return &NotFoundError{baseError: base}
		}

	case http.StatusBadRequest, http.StatusConflict, http.StatusTooManyRequests:
		var body struct {
			Message string `json:"message"`
		}
		if json.Unmarshal(data, &body) == nil && body.Message != "" {
			base.Message = body.Message
			return &GenericError{baseError: base}
		}
	}

	// Undocumented status, or a documented one whose body didn't match.
	base.Message = strings.TrimSpace(string(data))
	return &UnexpectedStatusError{baseError: base, Body: data}
}
