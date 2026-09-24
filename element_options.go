package tourmanager

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// PriceUnit determines the unit an ElementOption's price is quoted in.
type PriceUnit string

const (
	// PriceUnitUnit prices are fixed regardless of duration or number of
	// passengers.
	PriceUnitUnit PriceUnit = "unit"
	// PriceUnitPerson prices are multiplied by the number of passengers.
	PriceUnitPerson PriceUnit = "person"
	// PriceUnitUnitDay prices are fixed for passengers, multiplied by
	// duration in days.
	PriceUnitUnitDay PriceUnit = "unit_day"
	// PriceUnitPersonDay prices are multiplied by both the number of
	// passengers and duration in days.
	PriceUnitPersonDay PriceUnit = "person_day"
)

// Occupancy bounds the number of passengers that may be allocated to a
// single reservation of an ElementOption.
type Occupancy struct {
	// From is the minimum number of passengers required to book.
	From int `json:"from"`
	// To is the maximum number of passengers bookable per reservation.
	To int `json:"to"`
}

// PassengerSex constrains an ElementOption to passengers of a given sex.
type PassengerSex string

const (
	PassengerSexMale   PassengerSex = "m"
	PassengerSexFemale PassengerSex = "f"
)

// ElementOptionConstraints holds the additional booking constraints an
// ElementOption may declare. openapi.yml documents this object with
// additionalProperties: true; only its explicitly documented fields are
// represented here.
type ElementOptionConstraints struct {
	// MinAge is the minimum age of all passengers at time of travel
	// (inclusive), if set.
	MinAge *int `json:"min_age,omitempty"`
	// MaxAge is the maximum age of all passengers at time of travel
	// (inclusive), if set.
	MaxAge *int `json:"max_age,omitempty"`
	// PassengerSex, if non-empty, restricts booking to passengers of the
	// listed sexes.
	PassengerSex []PassengerSex `json:"passenger_sex,omitempty"`
}

// ElementOption is one bookable configuration of an [Element] (for example,
// a room type or a fare class). Every ElementOption on an Element shares
// that Element's inventory.
type ElementOption struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Name        string    `json:"name"`
	Category    Category  `json:"category"`
	PriceUnit   PriceUnit `json:"price_unit"`
	Occupancy   Occupancy `json:"occupancy"`
	Published   bool      `json:"published"`
	WebBookable bool      `json:"web_bookable"`
	IsLead      bool      `json:"is_lead"`

	Constraints *ElementOptionConstraints `json:"constraints,omitempty"`
}

// CreateElementOptionParams are the properties for creating a new
// ElementOption. Name, CategoryID, Occupancy and PriceUnit are required.
//
// spec: createOption
type CreateElementOptionParams struct {
	Name       string    `json:"name"`
	CategoryID string    `json:"category_id"`
	Occupancy  Occupancy `json:"occupancy"`
	PriceUnit  PriceUnit `json:"price_unit"`

	Published   *bool `json:"published,omitempty"`
	WebBookable *bool `json:"web_bookable,omitempty"`
	IsLead      *bool `json:"is_lead,omitempty"`
}

// UpdateElementOptionParams are the properties that may change on an
// existing ElementOption. All fields are optional; a nil/zero field is left
// unchanged.
//
// spec: updateOption
type UpdateElementOptionParams struct {
	Name       *string    `json:"name,omitempty"`
	CategoryID *string    `json:"category_id,omitempty"`
	Occupancy  *Occupancy `json:"occupancy,omitempty"`
	PriceUnit  PriceUnit  `json:"price_unit,omitempty"`

	Published   *bool `json:"published,omitempty"`
	WebBookable *bool `json:"web_bookable,omitempty"`
	IsLead      *bool `json:"is_lead,omitempty"`
}

// ElementOptions manages the Options belonging to one specific Element.
//
// Obtain one via [ElementsResource.Options]. Unlike [CategoryAttachment],
// this is not an attach/detach relationship — it's plain CRUD, following
// the same shape as [HolidayRelations]. openapi.yml also defines no list
// operation for an Element's Options: they're only discoverable nested
// inside [Element.Options], or inside a Departure's own tree (see
// [DepartureElementOption]).
type ElementOptions struct {
	client    *Client
	elementID string
}

func (o *ElementOptions) path(suffix string) string {
	return "/holidays/elements/" + url.PathEscape(o.elementID) + "/options" + suffix
}

// Create creates a new Option on this Element.
//
// spec: createOption
func (o *ElementOptions) Create(ctx context.Context, params CreateElementOptionParams) (*ElementOption, error) {
	var opt ElementOption
	if err := o.client.do(ctx, http.MethodPost, o.path(""), nil, params, &opt); err != nil {
		return nil, err
	}
	return &opt, nil
}

// Get returns a single Option by ID.
//
// spec: getOption
func (o *ElementOptions) Get(ctx context.Context, optionID string) (*ElementOption, error) {
	var opt ElementOption
	if err := o.client.do(ctx, http.MethodGet, o.path("/"+url.PathEscape(optionID)), nil, nil, &opt); err != nil {
		return nil, err
	}
	return &opt, nil
}

// Update partially updates an existing Option.
//
// spec: updateOption
func (o *ElementOptions) Update(ctx context.Context, optionID string, params UpdateElementOptionParams) (*ElementOption, error) {
	var opt ElementOption
	if err := o.client.do(ctx, http.MethodPatch, o.path("/"+url.PathEscape(optionID)), nil, params, &opt); err != nil {
		return nil, err
	}
	return &opt, nil
}

// Delete deletes an Option.
//
// spec: deleteOption
func (o *ElementOptions) Delete(ctx context.Context, optionID string) error {
	return o.client.do(ctx, http.MethodDelete, o.path("/"+url.PathEscape(optionID)), nil, nil, nil)
}

// Restore restores a previously deleted Option. Unlike most restore
// endpoints, this one returns the restored Option.
//
// spec: restoreAnOption
func (o *ElementOptions) Restore(ctx context.Context, optionID string) (*ElementOption, error) {
	return restoreInto[ElementOption](ctx, o.client, o.path("/"+url.PathEscape(optionID)+"/restore"))
}
