package tourmanager

import (
	"context"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// HolidaySortField is a field Holidays can be sorted by.
type HolidaySortField string

const (
	HolidaySortID        HolidaySortField = "id"
	HolidaySortCode      HolidaySortField = "code"
	HolidaySortName      HolidaySortField = "name"
	HolidaySortOrdering  HolidaySortField = "ordering"
	HolidaySortCreatedAt HolidaySortField = "created_at"
	HolidaySortUpdatedAt HolidaySortField = "updated_at"
	HolidaySortDeletedAt HolidaySortField = "deleted_at"
)

// ListHolidaysOptions filters, sorts and paginates a call to
// [HolidaysResource.List] or [HolidaysResource.All]. The zero value lists
// the first page using the API's defaults.
type ListHolidaysOptions struct {
	ListOptions

	Sort  HolidaySortField
	Order SortOrder

	// Search performs a free-text search.
	Search string
	// Name filters holidays whose name contains this value.
	Name string
	// Code filters holidays whose code is prefixed with this value.
	Code string

	// Published and Trash filter by publish/trash state. nil means "don't
	// filter"; a non-nil value filters to exactly that state.
	Published *bool
	Trash     *bool
}

func (o ListHolidaysOptions) values() url.Values {
	q := o.ListOptions.values()
	setEnum(q, "sort", o.Sort)
	setEnum(q, "order", o.Order)
	setString(q, "search", o.Search)
	setString(q, "name", o.Name)
	setString(q, "code", o.Code)
	setBool01(q, "published", o.Published)
	setBool01(q, "trash", o.Trash)
	return q
}

// Holiday is RezKit's top-level bookable touring product.
type Holiday struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Code         string     `json:"code"`
	Introduction string     `json:"introduction"`
	Description  string     `json:"description"`
	Ordering     int        `json:"ordering"`
	Published    bool       `json:"published"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`

	// Fields holds this holiday's custom field data. openapi.yml's Holiday
	// schema lists "fields" in its `required` array but never defines it
	// under `properties` (a spec bug, not a gap — the update request body
	// does reference CustomFieldsData for the same field, confirming its
	// existence).
	Fields CustomFieldsData `json:"fields,omitempty"`

	// Note: the Holiday schema's `required` list also includes "rank",
	// which is defined nowhere in `properties` and doesn't correspond to
	// any documented field (likely a leftover from a rename to
	// "ordering") — not represented here. `slug` and `seo` appear on the
	// hand-written JS client but have zero evidence in openapi.yml's
	// Holiday schema (neither `properties` nor `required`) — also not
	// represented here pending spec confirmation.
}

// CreateHolidayParams are the properties for creating a new Holiday. Name
// and Code are required; the API defaults Published to false if omitted.
//
// spec: components.requestBodies.CreateHoliday
type CreateHolidayParams struct {
	Name string `json:"name"`
	Code string `json:"code"`

	Introduction *string `json:"introduction,omitempty"`
	Description  *string `json:"description,omitempty"`
	Published    *bool   `json:"published,omitempty"`
}

// UpdateHolidayParams are the properties that may change on an existing
// Holiday. All fields are optional; a nil field is left unchanged.
//
// Introduction and Description are `nullable: true` in openapi.yml's
// UpdateHoliday request body, meaning the API distinguishes "don't touch
// this field" from "explicitly clear it to null" — a distinction a plain
// *string can't express. Use [NullValue] to set a value or [Null] to
// clear the field; leave the pointer nil to leave it unchanged.
//
// spec: components.requestBodies.UpdateHoliday
type UpdateHolidayParams struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`

	Introduction *Nullable[string] `json:"introduction,omitempty"`
	Description  *Nullable[string] `json:"description,omitempty"`

	// Fields replaces this holiday's custom field data.
	Fields CustomFieldsData `json:"fields,omitempty"`

	Published *bool            `json:"published,omitempty"`
	Ordering  *OrderingCommand `json:"ordering,omitempty"`
}

// HolidaysResource provides access to the Holidays API.
//
// Obtain one via [Client.Holidays].
type HolidaysResource struct {
	client *Client
}

// Holidays returns a resource handle for the Holidays API.
func (c *Client) Holidays() *HolidaysResource {
	return &HolidaysResource{client: c}
}

// List returns one page of Holidays matching opts.
//
// spec: listHolidays
func (r *HolidaysResource) List(ctx context.Context, opts ListHolidaysOptions) (*Page[Holiday], error) {
	return fetchPage[Holiday](ctx, r.client, http.MethodGet, "/holidays", opts.values())
}

// All returns an iterator that lazily fetches every page of Holidays
// matching opts, yielding one Holiday (or an error) at a time.
//
//	for holiday, err := range client.Holidays().All(ctx, tourmanager.ListHolidaysOptions{}) {
//		if err != nil {
//			return err
//		}
//		fmt.Println(holiday.Code)
//	}
//
// spec: listHolidays
func (r *HolidaysResource) All(ctx context.Context, opts ListHolidaysOptions) iter.Seq2[Holiday, error] {
	return autoPaginate(ctx, opts.Page, func(ctx context.Context, page int) (*Page[Holiday], error) {
		opts.Page = page
		return r.List(ctx, opts)
	})
}

// Get returns a single Holiday by ID.
//
// spec: getHoliday
func (r *HolidaysResource) Get(ctx context.Context, id string) (*Holiday, error) {
	var h Holiday
	if err := r.client.do(ctx, http.MethodGet, "/holidays/"+url.PathEscape(id), nil, nil, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

// Create creates a new Holiday.
//
// spec: createHoliday
func (r *HolidaysResource) Create(ctx context.Context, params CreateHolidayParams) (*Holiday, error) {
	var h Holiday
	if err := r.client.do(ctx, http.MethodPost, "/holidays", nil, params, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

// Update partially updates an existing Holiday.
//
// spec: updateHoliday
func (r *HolidaysResource) Update(ctx context.Context, id string, params UpdateHolidayParams) (*Holiday, error) {
	var h Holiday
	if err := r.client.do(ctx, http.MethodPatch, "/holidays/"+url.PathEscape(id), nil, params, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

// Delete moves a Holiday into the trash. Use Restore to undo.
//
// spec: deleteHoliday
func (r *HolidaysResource) Delete(ctx context.Context, id string) error {
	return r.client.do(ctx, http.MethodDelete, "/holidays/"+url.PathEscape(id), nil, nil, nil)
}

// Restore restores a previously deleted Holiday out of the trash. Unlike
// most restore endpoints, this one returns the restored Holiday.
//
// spec: restoreAHoliday
func (r *HolidaysResource) Restore(ctx context.Context, id string) (*Holiday, error) {
	return restoreInto[Holiday](ctx, r.client, "/holidays/"+url.PathEscape(id)+"/restore")
}

// Categories returns the Category-attachment handle for the given
// Holiday. Equivalent to
// client.Categories().For(tourmanager.EntityTypeHoliday, id).
func (r *HolidaysResource) Categories(id string) *CategoryAttachment {
	return r.client.Categories().For(EntityTypeHoliday, id)
}

// Relations returns the relations handle for the given Holiday.
func (r *HolidaysResource) Relations(id string) *HolidayRelations {
	return &HolidayRelations{client: r.client, holidayID: id}
}

// Versions returns the Versions handle for the given Holiday.
func (r *HolidaysResource) Versions(id string) *HolidayVersions {
	return &HolidayVersions{client: r.client, holidayID: id}
}

// Copy (PUT /holidays/{id}/copy) is not implemented: openapi.yml documents
// the response (a Holiday) but not the request body at all. Implement
// once the spec documents what CopyHolidayParams should contain.
