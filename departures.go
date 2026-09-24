package tourmanager

import (
	"context"
	"encoding/json"
	"iter"
	"net/http"
	"net/url"
	"time"
)

// DepartureRangeType describes how a Departure's date range behaves.
type DepartureRangeType string

const (
	DepartureRangeFixed         DepartureRangeType = "fixed"
	DepartureRangeFixedDuration DepartureRangeType = "fixed_duration"
	DepartureRangeFlexible      DepartureRangeType = "flexible"
)

// ListDeparturesOptions filters and paginates a call to
// [DeparturesResource.List] or [DeparturesResource.All]. The zero value
// lists the first page using the API's defaults.
//
// Note: openapi.yml's listDepartures operation does not document page/sort
// query parameters explicitly (unlike listHolidays/listCategories), even
// though its response has the same paginated envelope shape. Page/Limit
// are still sent here on the well-corroborated assumption that this
// endpoint honors the same page/limit convention used identically
// elsewhere in the spec — confirm against a live response if this proves
// wrong.
type ListDeparturesOptions struct {
	ListOptions

	// VersionID filters to departures of a specific holiday version.
	VersionID string
	// HolidayID filters to departures of a specific holiday.
	HolidayID string
	// Before returns only departures which may return before this time.
	Before *time.Time
	// After returns only departures which may depart after this time.
	After *time.Time
	// Published filters by publish state. nil means "don't filter".
	Published *bool
}

func (o ListDeparturesOptions) values() url.Values {
	q := o.ListOptions.values()
	setString(q, "version", o.VersionID)
	setString(q, "holiday", o.HolidayID)
	if o.Before != nil {
		q.Set("before", o.Before.Format(time.RFC3339))
	}
	if o.After != nil {
		q.Set("after", o.After.Format(time.RFC3339))
	}
	setBool01(q, "published", o.Published)
	return q
}

// Departure is a specific bookable instance (or date range) of a Holiday.
type Departure struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	VersionID string             `json:"version_id"`
	RangeType DepartureRangeType `json:"range_type"`
	Start     time.Time          `json:"start"`
	End       time.Time          `json:"end"`
	Published bool               `json:"published"`
	Inventory Inventory          `json:"inventory"`

	// SourceID is the ID of the departure this one was copied from, if
	// any.
	SourceID *string `json:"source_id,omitempty"`

	// Elements holds this Departure's configured Elements (each with its
	// own inventory, Options and generated Price IDs) as returned inline
	// by the API. This is the only way to discover the IDs Tour Manager
	// generates for a DepartureElement, its Options and their Prices —
	// there is no separate endpoint to list them.
	Elements []DepartureElement `json:"elements,omitempty"`
}

// UnmarshalJSON decodes a Departure, dispatching its Inventory field to
// the concrete type identified by that field's own "type" discriminator.
func (d *Departure) UnmarshalJSON(data []byte) error {
	type alias Departure
	aux := struct {
		Inventory json.RawMessage `json:"inventory"`
		*alias
	}{alias: (*alias)(d)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	inv, err := unmarshalInventory(aux.Inventory)
	if err != nil {
		return err
	}
	d.Inventory = inv
	return nil
}

// DepartureElement is one [Element] as configured on a specific Departure:
// its own Inventory (independent of the Element's DefaultInventory), a
// balance-due date, and the [DepartureElementOption]s (with their generated
// [Price] IDs) bookable against it.
//
// Note: openapi.yml documents no way to write a DepartureElement's
// Inventory — there is no path anywhere in the spec shaped like
// /holidays/departures/{departure}/elements/{element}. The JS sibling
// client having such a call is not, on its own, sufficient evidence to
// guess and ship a request shape here (see AGENTS.md and the
// update-tour-manager-go-client skill, section 4); this is tracked as a
// backlog item in README.md until openapi.yml documents the operation.
type DepartureElement struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Inventory  Inventory                   `json:"inventory"`
	Element    DepartureElementSummary     `json:"element"`
	BalanceDue *DepartureElementBalanceDue `json:"balance_due,omitempty"`
	Options    []DepartureElementOption    `json:"options,omitempty"`
}

// UnmarshalJSON decodes a DepartureElement, dispatching its Inventory field
// to the concrete type identified by that field's own "type" discriminator.
func (e *DepartureElement) UnmarshalJSON(data []byte) error {
	type alias DepartureElement
	aux := struct {
		Inventory json.RawMessage `json:"inventory"`
		*alias
	}{alias: (*alias)(e)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	inv, err := unmarshalInventory(aux.Inventory)
	if err != nil {
		return err
	}
	e.Inventory = inv
	return nil
}

// DepartureElementSummary is the trimmed [Element] reference nested inside
// a [DepartureElement].
type DepartureElementSummary struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsPackage bool   `json:"is_package"`
	Published bool   `json:"published"`
}

// DepartureElementBalanceDue is the balance payment due date for a
// DepartureElement.
type DepartureElementBalanceDue struct {
	// Calculated reports whether Date was derived from a global rule, as
	// opposed to being set explicitly.
	Calculated bool      `json:"calculated"`
	Date       time.Time `json:"date"`
}

// DepartureElementOption is an [ElementOption] as offered on a specific
// Departure, together with its generated [Price]s — exactly one per
// operator currency.
type DepartureElementOption struct {
	ElementOption
	Prices []Price `json:"prices"`
}

// CreateDepartureParams are the properties for creating a new Departure.
// Start, End and Inventory are required.
//
// spec: DepartureProperties (allOf DepartureParams + its own `required`)
type CreateDepartureParams struct {
	Start     time.Time          `json:"start"`
	End       time.Time          `json:"end"`
	Inventory Inventory          `json:"inventory"`
	RangeType DepartureRangeType `json:"range_type,omitempty"`
}

// UpdateDepartureParams are the properties that may change on an existing
// Departure. All fields are optional; a nil/zero field is left unchanged.
//
// spec: DepartureParams
type UpdateDepartureParams struct {
	Start     *time.Time         `json:"start,omitempty"`
	End       *time.Time         `json:"end,omitempty"`
	Inventory Inventory          `json:"inventory,omitempty"`
	RangeType DepartureRangeType `json:"range_type,omitempty"`
}

// DeparturesResource provides access to the Departures API.
//
// Obtain one via [Client.Departures].
type DeparturesResource struct {
	client *Client
}

// Departures returns a resource handle for the Departures API.
func (c *Client) Departures() *DeparturesResource {
	return &DeparturesResource{client: c}
}

// List returns one page of Departures matching opts.
//
// spec: listDepartures
func (r *DeparturesResource) List(ctx context.Context, opts ListDeparturesOptions) (*Page[Departure], error) {
	return fetchPage[Departure](ctx, r.client, http.MethodGet, "/holidays/departures", opts.values())
}

// All auto-paginates List.
//
// spec: listDepartures
func (r *DeparturesResource) All(ctx context.Context, opts ListDeparturesOptions) iter.Seq2[Departure, error] {
	return autoPaginate(ctx, opts.Page, func(ctx context.Context, page int) (*Page[Departure], error) {
		opts.Page = page
		return r.List(ctx, opts)
	})
}

// Get returns a single Departure by ID.
//
// spec: getDeparture
func (r *DeparturesResource) Get(ctx context.Context, id string) (*Departure, error) {
	var d Departure
	if err := r.client.do(ctx, http.MethodGet, "/holidays/departures/"+url.PathEscape(id), nil, nil, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// Create creates a new Departure.
//
// spec: createDeparture
func (r *DeparturesResource) Create(ctx context.Context, params CreateDepartureParams) (*Departure, error) {
	var d Departure
	if err := r.client.do(ctx, http.MethodPost, "/holidays/departures", nil, params, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// Update partially updates an existing Departure.
//
// spec: updateDeparture
func (r *DeparturesResource) Update(ctx context.Context, id string, params UpdateDepartureParams) (*Departure, error) {
	var d Departure
	path := "/holidays/departures/" + url.PathEscape(id)
	if err := r.client.do(ctx, http.MethodPatch, path, nil, params, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// Delete deletes a Departure.
//
// Note: openapi.yml defines no restore operation for Departures (unlike
// most other resources), so no Restore method is exposed here.
//
// spec: deleteDeparture
func (r *DeparturesResource) Delete(ctx context.Context, id string) error {
	return r.client.do(ctx, http.MethodDelete, "/holidays/departures/"+url.PathEscape(id), nil, nil, nil)
}
