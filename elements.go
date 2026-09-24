package tourmanager

import (
	"context"
	"encoding/json"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// ElementSortField is a field Elements can be sorted by.
type ElementSortField string

const (
	ElementSortID        ElementSortField = "id"
	ElementSortName      ElementSortField = "name"
	ElementSortCreatedAt ElementSortField = "created_at"
	ElementSortUpdatedAt ElementSortField = "updated_at"
)

// ListElementsOptions filters, sorts and paginates a call to
// [ElementsResource.List] or [ElementsResource.All]. The zero value lists
// the first page using the API's defaults.
type ListElementsOptions struct {
	ListOptions

	Sort  ElementSortField
	Order SortOrder

	// Name filters elements whose name contains this value.
	Name string

	// Trash and Published filter by trash/publish state. nil means "don't
	// filter"; a non-nil value filters to exactly that state.
	Trash     *bool
	Published *bool
}

func (o ListElementsOptions) values() url.Values {
	q := o.ListOptions.values()
	setEnum(q, "sort", o.Sort)
	setEnum(q, "order", o.Order)
	setString(q, "name", o.Name)
	setBool01(q, "trash", o.Trash)
	setBool01(q, "published", o.Published)
	return q
}

// Element is a bookable component of a Holiday Version (e.g. accommodation,
// transport, or an activity), offered as one or more [ElementOption]s that
// all share the Element's inventory.
type Element struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	VersionID        string    `json:"version_id"`
	Name             string    `json:"name"`
	Category         Category  `json:"category"`
	IsPackage        bool      `json:"is_package"`
	Published        bool      `json:"published"`
	DefaultInventory Inventory `json:"default_inventory"`

	// BalanceDue is a rule for balance payment due relative to the
	// departure date. nil means no element-specific rule is set and the
	// operator's global default applies.
	BalanceDue *int `json:"balance_due,omitempty"`

	// Options holds this Element's available Options, as returned inline
	// by the API. Every Option on an Element shares that Element's
	// inventory.
	Options []ElementOption `json:"options,omitempty"`
}

// UnmarshalJSON decodes an Element, dispatching its DefaultInventory field
// to the concrete type identified by that field's own "type" discriminator.
func (e *Element) UnmarshalJSON(data []byte) error {
	type alias Element
	aux := struct {
		DefaultInventory json.RawMessage `json:"default_inventory"`
		*alias
	}{alias: (*alias)(e)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	inv, err := unmarshalInventory(aux.DefaultInventory)
	if err != nil {
		return err
	}
	e.DefaultInventory = inv
	return nil
}

// CreateElementParams are the properties for creating a new Element. Name
// and CategoryID are required.
//
// spec: createElement
type CreateElementParams struct {
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`

	DefaultInventory Inventory `json:"default_inventory,omitempty"`
	Published        *bool     `json:"published,omitempty"`
	IsPackage        *bool     `json:"is_package,omitempty"`

	// BalanceDue is `nullable: true` in openapi.yml's create body (as it is
	// on update): use [NullValue] to set it, [Null] to explicitly send
	// null, or leave this nil to omit the field entirely.
	BalanceDue *Nullable[int] `json:"balance_due,omitempty"`
}

// UpdateElementParams are the properties that may change on an existing
// Element. All fields are optional; a nil field is left unchanged.
//
// spec: updateElement
type UpdateElementParams struct {
	Name       *string `json:"name,omitempty"`
	CategoryID *string `json:"category_id,omitempty"`

	DefaultInventory Inventory `json:"default_inventory,omitempty"`
	Published        *bool     `json:"published,omitempty"`
	IsPackage        *bool     `json:"is_package,omitempty"`

	// BalanceDue is `nullable: true`; see [CreateElementParams.BalanceDue].
	BalanceDue *Nullable[int] `json:"balance_due,omitempty"`
}

// ElementsResource provides access to the Elements API.
//
// Obtain one via [Client.Elements].
type ElementsResource struct {
	client *Client
}

// Elements returns a resource handle for the Elements API.
func (c *Client) Elements() *ElementsResource {
	return &ElementsResource{client: c}
}

func elementPath(versionID, suffix string) string {
	return "/holidays/versions/" + url.PathEscape(versionID) + "/elements" + suffix
}

// List returns one page of Elements belonging to the given Holiday Version.
//
// spec: listElements
func (r *ElementsResource) List(ctx context.Context, versionID string, opts ListElementsOptions) (*Page[Element], error) {
	return fetchPage[Element](ctx, r.client, http.MethodGet, elementPath(versionID, ""), opts.values())
}

// All auto-paginates List.
//
// spec: listElements
func (r *ElementsResource) All(ctx context.Context, versionID string, opts ListElementsOptions) iter.Seq2[Element, error] {
	return autoPaginate(ctx, opts.Page, func(ctx context.Context, page int) (*Page[Element], error) {
		opts.Page = page
		return r.List(ctx, versionID, opts)
	})
}

// Get returns a single Element by ID.
//
// spec: getElement
func (r *ElementsResource) Get(ctx context.Context, versionID, elementID string) (*Element, error) {
	var e Element
	path := elementPath(versionID, "/"+url.PathEscape(elementID))
	if err := r.client.do(ctx, http.MethodGet, path, nil, nil, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

// Create creates a new Element on the given Holiday Version.
//
// spec: createElement
func (r *ElementsResource) Create(ctx context.Context, versionID string, params CreateElementParams) (*Element, error) {
	var e Element
	if err := r.client.do(ctx, http.MethodPost, elementPath(versionID, ""), nil, params, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

// Update partially updates an existing Element.
//
// spec: updateElement
func (r *ElementsResource) Update(ctx context.Context, versionID, elementID string, params UpdateElementParams) (*Element, error) {
	var e Element
	path := elementPath(versionID, "/"+url.PathEscape(elementID))
	if err := r.client.do(ctx, http.MethodPatch, path, nil, params, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

// Delete deletes an Element.
//
// Note: openapi.yml defines no restore operation for Elements (unlike most
// other resources), so no Restore method is exposed here.
//
// spec: deleteElement
func (r *ElementsResource) Delete(ctx context.Context, versionID, elementID string) error {
	path := elementPath(versionID, "/"+url.PathEscape(elementID))
	return r.client.do(ctx, http.MethodDelete, path, nil, nil, nil)
}

// Options returns the ElementOptions handle for the given Element.
func (r *ElementsResource) Options(elementID string) *ElementOptions {
	return &ElementOptions{client: r.client, elementID: elementID}
}

// Categories returns the Category-attachment handle for the given Element.
// Equivalent to client.Categories().For(tourmanager.EntityTypeElement, id).
func (r *ElementsResource) Categories(id string) *CategoryAttachment {
	return r.client.Categories().For(EntityTypeElement, id)
}
