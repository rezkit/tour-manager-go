package tourmanager

import (
	"context"
	"net/http"
)

// restoreEmpty issues a PUT request to path and discards the response
// body — used by the majority of restore endpoints, which return 202 with
// no response schema.
func restoreEmpty(ctx context.Context, c *Client, path string) error {
	return c.do(ctx, http.MethodPut, path, nil, nil, nil)
}

// restoreInto issues a PUT request to path and decodes the restored
// resource from the response — used by the minority of restore endpoints
// that return the resource in their response body.
func restoreInto[T any](ctx context.Context, c *Client, path string) (*T, error) {
	var v T
	if err := c.do(ctx, http.MethodPut, path, nil, nil, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
