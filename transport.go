package tourmanager

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// do executes an HTTP request against the API and is the single entry
// point every resource method funnels through.
//
//   - path is relative to the client's base URL and must start with "/".
//     Callers are responsible for url.PathEscape-ing any caller-supplied
//     value (an ID, an EntityType, ...) interpolated into it.
//   - query, if non-nil, is appended as the request's query string.
//   - body, if non-nil, is JSON-encoded as the request body.
//   - out, if non-nil, receives the JSON-decoded response body on success.
//
// A non-2xx response becomes one of the [APIError] implementations in
// errors.go. A request that never received a response becomes a
// [RequestError]; a response whose body didn't decode as expected becomes
// a [DecodeError].
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body, out any) error {
	req, err := c.newRequest(ctx, method, path, query, body)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &RequestError{Method: method, URL: req.URL.String(), Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return newAPIError(method, req.URL.String(), resp)
	}

	if out == nil || resp.StatusCode == http.StatusNoContent {
		_, _ = io.Copy(io.Discard, resp.Body) // drain for connection reuse
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return &DecodeError{Method: method, URL: req.URL.String(), Err: err}
	}
	return nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, query url.Values, body any) (*http.Request, error) {
	u, err := c.buildURL(path, query)
	if err != nil {
		return nil, err
	}

	var bodyReader io.Reader
	if body != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, fmt.Errorf("tourmanager: encode request body: %w", err)
		}
		bodyReader = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("tourmanager: build request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", c.userAgent)

	return req, nil
}

// buildURL joins relPath onto the client's base URL.
func (c *Client) buildURL(relPath string, query url.Values) (*url.URL, error) {
	base := c.baseURL
	if base == "" {
		base = defaultBaseURL
	}
	u, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("tourmanager: invalid base URL %q: %w", base, err)
	}
	u.Path = strings.TrimRight(u.Path, "/") + "/" + strings.TrimLeft(relPath, "/")
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	return u, nil
}
