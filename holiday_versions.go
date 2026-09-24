package tourmanager

import (
	"context"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// HolidayVersionSortField is a field Holiday Versions can be sorted by.
//
// Note: openapi.yml's listHolidayVersions operation reuses the exact same
// sort parameter as listHolidays, whose enum includes "ordering" and
// "deleted_at" — HolidayVersion's own schema exposes neither field. Kept
// exactly as documented since this is a query parameter, not a payload
// shape the response needs to satisfy.
type HolidayVersionSortField string

const (
	HolidayVersionSortID        HolidayVersionSortField = "id"
	HolidayVersionSortCode      HolidayVersionSortField = "code"
	HolidayVersionSortName      HolidayVersionSortField = "name"
	HolidayVersionSortOrdering  HolidayVersionSortField = "ordering"
	HolidayVersionSortCreatedAt HolidayVersionSortField = "created_at"
	HolidayVersionSortUpdatedAt HolidayVersionSortField = "updated_at"
	HolidayVersionSortDeletedAt HolidayVersionSortField = "deleted_at"
)

// ListHolidayVersionsOptions filters, sorts and paginates a call to
// [HolidayVersions.List] or [HolidayVersions.All]. The zero value lists the
// first page using the API's defaults.
type ListHolidayVersionsOptions struct {
	ListOptions

	Sort  HolidayVersionSortField
	Order SortOrder

	// Search performs a free-text search.
	Search string
	// Name filters versions whose name contains this value.
	Name string
	// Code filters versions whose code is prefixed with this value.
	Code string

	// Published and Trash filter by publish/trash state. nil means "don't
	// filter"; a non-nil value filters to exactly that state.
	Published *bool
	Trash     *bool
}

func (o ListHolidayVersionsOptions) values() url.Values {
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

// HolidayVersion is one dated/priced configuration of a Holiday — the unit
// that Elements, Departures and itinerary entries actually attach to.
//
// Note: unlike most other resources in this client, openapi.yml's
// HolidayVersion schema has no `deleted_at` property (and no `required`
// list at all), even though deleteHolidayVersion/restoreHolidayVersion
// both exist — not represented here pending spec confirmation.
type HolidayVersion struct {
	ID        string    `json:"id"`
	HolidayID string    `json:"holiday_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name         string `json:"name"`
	Code         string `json:"code"`
	Introduction string `json:"introduction"`
	Description  string `json:"description"`
	Published    bool   `json:"published"`
}

// CreateHolidayVersionParams are the properties for creating a new
// HolidayVersion. Name and Code are required.
//
// openapi.yml's createHolidayVersion operation reuses the same request
// body as createHoliday (components.requestBodies.CreateHoliday) — every
// field in it also exists on HolidayVersion itself, so it's imposed here
// as its own parallel type rather than sharing CreateHolidayParams,
// matching this client's usual one-Params-type-per-resource convention.
//
// spec: createHolidayVersion
type CreateHolidayVersionParams struct {
	Name string `json:"name"`
	Code string `json:"code"`

	Introduction *string `json:"introduction,omitempty"`
	Description  *string `json:"description,omitempty"`
	Published    *bool   `json:"published,omitempty"`
}

// UpdateHolidayVersionParams are the properties that may change on an
// existing HolidayVersion. All fields are optional; a nil field is left
// unchanged.
//
// Introduction and Description are `nullable: true`, as on
// [UpdateHolidayParams] — use [NullValue] to set a value or [Null] to
// clear it; leave the pointer nil to leave it unchanged.
//
// Note: openapi.yml's updateHolidayVersion operation reuses
// components.requestBodies.UpdateHoliday verbatim, which also documents
// "fields" (CustomFieldsData) and "ordering" (OrderingCommand) properties.
// Neither has any corresponding property on HolidayVersion's own read
// schema, unlike Name/Code/Introduction/Description/Published — treated as
// artifacts of reusing Holiday's update body rather than confirmed
// HolidayVersion capabilities, and excluded here pending spec
// clarification (same reasoning as Holiday's excluded slug/seo/rank
// fields — see README.md's "Known gaps").
//
// spec: updateHolidayVersion
type UpdateHolidayVersionParams struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`

	Introduction *Nullable[string] `json:"introduction,omitempty"`
	Description  *Nullable[string] `json:"description,omitempty"`

	Published *bool `json:"published,omitempty"`
}

// HolidayVersions manages the Versions belonging to one specific Holiday.
//
// Obtain one via [HolidaysResource.Versions].
type HolidayVersions struct {
	client    *Client
	holidayID string
}

func (r *HolidayVersions) path(suffix string) string {
	return "/holidays/" + url.PathEscape(r.holidayID) + "/versions" + suffix
}

// List returns one page of Versions belonging to this Holiday.
//
// spec: listHolidayVersions
func (r *HolidayVersions) List(ctx context.Context, opts ListHolidayVersionsOptions) (*Page[HolidayVersion], error) {
	return fetchPage[HolidayVersion](ctx, r.client, http.MethodGet, r.path(""), opts.values())
}

// All auto-paginates List.
//
// spec: listHolidayVersions
func (r *HolidayVersions) All(ctx context.Context, opts ListHolidayVersionsOptions) iter.Seq2[HolidayVersion, error] {
	return autoPaginate(ctx, opts.Page, func(ctx context.Context, page int) (*Page[HolidayVersion], error) {
		opts.Page = page
		return r.List(ctx, opts)
	})
}

// Get returns a single Version by ID.
//
// spec: getHolidayVersion
func (r *HolidayVersions) Get(ctx context.Context, versionID string) (*HolidayVersion, error) {
	var v HolidayVersion
	if err := r.client.do(ctx, http.MethodGet, r.path("/"+url.PathEscape(versionID)), nil, nil, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// Create creates a new Version on this Holiday.
//
// spec: createHolidayVersion
func (r *HolidayVersions) Create(ctx context.Context, params CreateHolidayVersionParams) (*HolidayVersion, error) {
	var v HolidayVersion
	if err := r.client.do(ctx, http.MethodPost, r.path(""), nil, params, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// Update partially updates an existing Version.
//
// spec: updateHolidayVersion
func (r *HolidayVersions) Update(ctx context.Context, versionID string, params UpdateHolidayVersionParams) (*HolidayVersion, error) {
	var v HolidayVersion
	path := r.path("/" + url.PathEscape(versionID))
	if err := r.client.do(ctx, http.MethodPatch, path, nil, params, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// Delete deletes a Version. Use Restore to undo.
//
// spec: deleteHolidayVersion
func (r *HolidayVersions) Delete(ctx context.Context, versionID string) error {
	return r.client.do(ctx, http.MethodDelete, r.path("/"+url.PathEscape(versionID)), nil, nil, nil)
}

// Restore restores a previously deleted Version. Unlike most restore
// endpoints, this one returns the restored Version.
//
// spec: restoreHolidayVersion
func (r *HolidayVersions) Restore(ctx context.Context, versionID string) (*HolidayVersion, error) {
	return restoreInto[HolidayVersion](ctx, r.client, r.path("/"+url.PathEscape(versionID)+"/restore"))
}
