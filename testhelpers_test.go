package tourmanager

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// newTestClient starts an httptest.Server using handler and returns a
// Client pointed at it. The server is closed automatically when t's test
// finishes.
func newTestClient(t *testing.T, handler http.Handler) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return New("test-api-key", WithBaseURL(srv.URL))
}

// handlerFunc adapts a plain function to http.Handler, for tests that only
// need to handle a single request shape.
type handlerFunc func(w http.ResponseWriter, r *http.Request)

func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) { f(w, r) }
