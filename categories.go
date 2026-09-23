package tourmanager

import (
	"context"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// Category is a named grouping that can be attached to entities of
// various [EntityType]s (holidays, elements, accommodations, ...) to
// classify or filter them.
type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// ParentID is the ID of this category's parent, if it has one.
	ParentID   string    `json:"parent_id,omitempty"`
	Published  bool      `json:"published"`
	Searchable bool      `json:"searchable"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Ancestors and Children are populated by the API only when the
	// corresponding list options (ListCategoriesOptions.Children) are
	// requested; otherwise they are empty.
	Ancestors []Category `json:"ancestors,omitempty"`
	Children  []Category `json:"children,omitempty"`
}

// CategorySortField is a field Categories can be sorted by.
type CategorySortField string

const (
	CategorySortID        CategorySortField = "id"
	CategorySortName      CategorySortField = "name"
	CategorySortParentID  CategorySortField = "parent_id"
	CategorySortCreatedAt CategorySortField = "created_at"
	CategorySortUpdatedAt CategorySortField = "updated_at"
)

// ListCategoriesOptions filters, sorts and paginates a call to
// [CategoriesResource.List]/[CategoriesResource.All] or
// [CategoryAttachment.List]/[CategoryAttachment.All]. The zero value lists
// the first page using the API's defaults.
type ListCategoriesOptions struct {
	ListOptions

	Sort  CategorySortField
	Order SortOrder

	// Name filters by category name.
	Name string
	// Search performs a free-text search.
	Search string
	// ParentID filters to categories with this parent.
	ParentID string

	// Children includes each category's child categories in the response.
	Children bool

	// Published and Searchable filter by publish/searchable state. nil
	// means "don't filter"; a non-nil value filters to exactly that state.
	Published  *bool
	Searchable *bool
}

func (o ListCategoriesOptions) values() url.Values {
	q := o.ListOptions.values()
	setEnum(q, "sort", o.Sort)
	setEnum(q, "order", o.Order)
	setString(q, "name", o.Name)
	setString(q, "search", o.Search)
	setString(q, "parent_id", o.ParentID)
	setBool(q, "children", o.Children)
	setBool01(q, "published", o.Published)
	setBool01(q, "searchable", o.Searchable)
	return q
}

// CreateCategoryParams are the properties for creating a new Category.
// openapi.yml defines this operation's request body as an untyped inline
// object; this type imposes the project's usual clean create/update
// convention over it.
//
// Name is required; all other fields are optional.
type CreateCategoryParams struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
	Published   *bool   `json:"published,omitempty"`
	Searchable  *bool   `json:"searchable,omitempty"`
}

// UpdateCategoryParams are the properties that may change on an existing
// Category. All fields are optional; a nil field is left unchanged.
//
// openapi.yml's inline PATCH body for categories has no "ordering" field
// (unlike, for example, Holiday's update body) — there is no
// Category-reorder support in this client until the spec documents one.
type UpdateCategoryParams struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
	Published   *bool   `json:"published,omitempty"`
	Searchable  *bool   `json:"searchable,omitempty"`
}

// CategoryAttachment manages the Categories attached to one specific item:
// an entity of some [EntityType] with a given ID.
//
// Obtain one via [CategoriesResource.For], or via a resource's own
// convenience accessor, such as [HolidaysResource.Categories].
type CategoryAttachment struct {
	h *AttachmentHandle[Category]
}

// List returns one page of Categories attached to this item.
//
// spec: listAttachedCategories
func (a *CategoryAttachment) List(ctx context.Context, opts ListCategoriesOptions) (*Page[Category], error) {
	return a.h.List(ctx, opts.values())
}

// All auto-paginates every Category attached to this item.
//
// spec: listAttachedCategories
func (a *CategoryAttachment) All(ctx context.Context, opts ListCategoriesOptions) iter.Seq2[Category, error] {
	return a.h.All(ctx, opts.Page, opts.values())
}

// Attach adds the given Category IDs to this item, preserving whatever is
// already attached. It returns the full set of Categories now attached.
//
// spec: attachCategories
func (a *CategoryAttachment) Attach(ctx context.Context, categoryIDs []string) ([]Category, error) {
	return a.h.Attach(ctx, categoryIDs)
}

// Replace overwrites this item's attached Categories with exactly
// categoryIDs.
//
// spec: replaceCategories
func (a *CategoryAttachment) Replace(ctx context.Context, categoryIDs []string) ([]Category, error) {
	return a.h.Replace(ctx, categoryIDs)
}

// Detach removes the given Category IDs from this item.
//
// spec: detachCategories
func (a *CategoryAttachment) Detach(ctx context.Context, categoryIDs []string) error {
	return a.h.Detach(ctx, categoryIDs)
}

// CategoriesResource provides access to the Categories API: both the
// parent, EntityType-scoped catalog of Category records, and, via
// [CategoriesResource.For], the per-item attachment sub-resource.
//
// Obtain one via [Client.Categories].
type CategoriesResource struct {
	client *Client
}

// Categories returns a resource handle for the Categories API.
func (c *Client) Categories() *CategoriesResource {
	return &CategoriesResource{client: c}
}

// For returns the attachment handle for the Categories attached to a
// specific item of the given EntityType.
//
//	cats, err := client.Categories().For(tourmanager.EntityTypeHoliday, holidayID).
//		List(ctx, tourmanager.ListCategoriesOptions{})
func (r *CategoriesResource) For(entity EntityType, itemID string) *CategoryAttachment {
	return &CategoryAttachment{h: newAttachmentHandle[Category](r.client, entity, itemID, "categories")}
}

// List returns one page of Category records belonging to entity's
// namespace (the parent catalog of categories, e.g. every "holiday"
// category — not the subset attached to a particular item; see
// [CategoriesResource.For] for that).
//
// spec: listCategories
func (r *CategoriesResource) List(ctx context.Context, entity EntityType, opts ListCategoriesOptions) (*Page[Category], error) {
	return fetchPage[Category](ctx, r.client, http.MethodGet, "/"+string(entity)+"/categories", opts.values())
}

// All auto-paginates List.
//
// spec: listCategories
func (r *CategoriesResource) All(ctx context.Context, entity EntityType, opts ListCategoriesOptions) iter.Seq2[Category, error] {
	return autoPaginate(ctx, opts.Page, func(ctx context.Context, page int) (*Page[Category], error) {
		opts.Page = page
		return r.List(ctx, entity, opts)
	})
}

// Create creates a new Category in entity's namespace.
//
// spec: createCategory
func (r *CategoriesResource) Create(ctx context.Context, entity EntityType, params CreateCategoryParams) (*Category, error) {
	var cat Category
	if err := r.client.do(ctx, http.MethodPost, "/"+string(entity)+"/categories", nil, params, &cat); err != nil {
		return nil, err
	}
	return &cat, nil
}

// Update changes properties of an existing Category.
//
// Note: openapi.yml does not define a single-category GET operation (only
// PATCH/DELETE exist at this path), so no Get method is exposed here.
//
// spec: updateCategory
func (r *CategoriesResource) Update(ctx context.Context, entity EntityType, id string, params UpdateCategoryParams) (*Category, error) {
	var cat Category
	path := "/" + string(entity) + "/categories/" + url.PathEscape(id)
	if err := r.client.do(ctx, http.MethodPatch, path, nil, params, &cat); err != nil {
		return nil, err
	}
	return &cat, nil
}

// Delete deletes a Category.
//
// spec: deleteCategory
func (r *CategoriesResource) Delete(ctx context.Context, entity EntityType, id string) error {
	path := "/" + string(entity) + "/categories/" + url.PathEscape(id)
	return r.client.do(ctx, http.MethodDelete, path, nil, nil, nil)
}

// Restore restores a previously deleted Category. The API returns 202
// with no response body for this endpoint.
//
// spec: restoreCategory
func (r *CategoriesResource) Restore(ctx context.Context, entity EntityType, id string) error {
	path := "/" + string(entity) + "/categories/" + url.PathEscape(id) + "/restore"
	return restoreEmpty(ctx, r.client, path)
}
