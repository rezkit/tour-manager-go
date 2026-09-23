package tourmanager

import (
	"context"
	"iter"
	"net/url"
)

// Page is one page of results from a paginated list endpoint, matching the
// API's pagination envelope.
type Page[T any] struct {
	// Total is the total number of items matching the request, across all
	// pages.
	Total int `json:"total"`
	// CurrentPage is the 1-based page number of this page.
	CurrentPage int `json:"current_page"`
	// LastPage is the number of the final page.
	LastPage int `json:"last_page"`
	// From is the index of the first item on this page.
	From int `json:"from"`
	// To is the index of the last item on this page.
	To int `json:"to"`
	// Data holds the items on this page.
	Data []T `json:"data"`
}

// ListOptions holds the pagination parameters shared by every paginated
// list endpoint. Resource-specific options types embed it:
//
//	type ListHolidaysOptions struct {
//	    ListOptions
//	    Sort  HolidaySortField
//	    Order SortOrder
//	}
//
// so callers can write &ListHolidaysOptions{Page: 2} directly via Go's
// field promotion.
type ListOptions struct {
	// Page is the 1-based page to fetch. Zero uses the API's default (1).
	Page int
	// Limit is the maximum number of items per page. Zero uses the API's
	// default.
	Limit int
}

func (o ListOptions) values() url.Values {
	v := url.Values{}
	setInt(v, "page", o.Page)
	setInt(v, "limit", o.Limit)
	return v
}

// listResponse is the shape of list endpoints that return only
// {"data": [...]}, with no pagination metadata.
type listResponse[T any] struct {
	Data []T `json:"data"`
}

// fetchPage performs a single paginated list request and decodes the
// response into a Page[T].
func fetchPage[T any](ctx context.Context, c *Client, method, path string, query url.Values) (*Page[T], error) {
	var page Page[T]
	if err := c.do(ctx, method, path, query, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// fetchList performs a single list request against an endpoint with no
// pagination metadata, decoding the response's "data" array.
func fetchList[T any](ctx context.Context, c *Client, method, path string, query url.Values) ([]T, error) {
	var resp listResponse[T]
	if err := c.do(ctx, method, path, query, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// autoPaginate returns an iterator over every item across every page that
// fetch produces, starting at startPage (values < 1 are treated as 1). It
// calls fetch again after exhausting each page until the page reports
// CurrentPage >= LastPage, or a page comes back with no items.
//
// If fetch returns an error, the iterator yields the zero value of T and
// that error once, then stops — a failure part-way through iteration is
// never silently dropped.
func autoPaginate[T any](ctx context.Context, startPage int, fetch func(ctx context.Context, page int) (*Page[T], error)) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		page := startPage
		if page < 1 {
			page = 1
		}
		for {
			result, err := fetch(ctx, page)
			if err != nil {
				var zero T
				yield(zero, err)
				return
			}
			for _, item := range result.Data {
				if !yield(item, nil) {
					return
				}
			}
			if len(result.Data) == 0 || result.CurrentPage >= result.LastPage {
				return
			}
			page = result.CurrentPage + 1
		}
	}
}
