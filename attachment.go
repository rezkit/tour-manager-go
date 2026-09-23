package tourmanager

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"
)

// AttachmentHandle provides list/attach/replace/detach access to items of
// type T attached to a single parent entity, following the convention used
// throughout the API:
//
//	GET    /{entity}/{id}/{segment}                -> paginated list of T
//	PATCH  /{entity}/{id}/{segment}  {ids: [...]}   -> attach (union), returns []T
//	PUT    /{entity}/{id}/{segment}  {ids: [...]}   -> replace (overwrite), returns []T
//	DELETE /{entity}/{id}/{segment}?ids=...         -> detach, 204 no content
//
// AttachmentHandle is unexported at the type-parameter level: resources
// wrap it in a small typed struct (see [CategoryAttachment]) so its
// methods render as concrete signatures in generated documentation,
// instead of a generic instantiation.
type AttachmentHandle[T any] struct {
	client  *Client
	entity  EntityType
	id      string
	segment string
}

// newAttachmentHandle constructs an AttachmentHandle scoped to a single
// parent entity and attachment path segment (e.g. "categories").
func newAttachmentHandle[T any](c *Client, entity EntityType, id, segment string) *AttachmentHandle[T] {
	return &AttachmentHandle[T]{client: c, entity: entity, id: id, segment: segment}
}

func (a *AttachmentHandle[T]) path() string {
	return fmt.Sprintf("/%s/%s/%s", a.entity, url.PathEscape(a.id), a.segment)
}

// List returns one page of items of type T attached to the parent entity,
// using query as additional filter/sort/pagination parameters.
func (a *AttachmentHandle[T]) List(ctx context.Context, query url.Values) (*Page[T], error) {
	return fetchPage[T](ctx, a.client, http.MethodGet, a.path(), query)
}

// All auto-paginates over every item of type T attached to the parent
// entity matching query, starting at startPage (see [autoPaginate]).
func (a *AttachmentHandle[T]) All(ctx context.Context, startPage int, query url.Values) iter.Seq2[T, error] {
	return autoPaginate(ctx, startPage, func(ctx context.Context, page int) (*Page[T], error) {
		q := cloneValues(query)
		setInt(q, "page", page)
		return a.List(ctx, q)
	})
}

type attachIDsBody struct {
	IDs []string `json:"ids"`
}

// Attach adds ids to the parent entity's existing attachment set,
// preserving whatever is already attached. It returns the full set of T
// now attached.
func (a *AttachmentHandle[T]) Attach(ctx context.Context, ids []string) ([]T, error) {
	var out []T
	err := a.client.do(ctx, http.MethodPatch, a.path(), nil, attachIDsBody{IDs: ids}, &out)
	return out, err
}

// Replace overwrites the parent entity's attachment set with exactly ids.
func (a *AttachmentHandle[T]) Replace(ctx context.Context, ids []string) ([]T, error) {
	var out []T
	err := a.client.do(ctx, http.MethodPut, a.path(), nil, attachIDsBody{IDs: ids}, &out)
	return out, err
}

// Detach removes ids from the parent entity's attachment set. ids are sent
// as a repeated query parameter, not a request body.
func (a *AttachmentHandle[T]) Detach(ctx context.Context, ids []string) error {
	q := url.Values{"ids": ids}
	return a.client.do(ctx, http.MethodDelete, a.path(), q, nil, nil)
}

func cloneValues(v url.Values) url.Values {
	out := make(url.Values, len(v))
	for k, vv := range v {
		out[k] = append([]string(nil), vv...)
	}
	return out
}
