package tourmanager

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// PriceDeposit describes the deposit payment required to confirm a
// reservation against a Price.
type PriceDeposit struct {
	// Calculated reports whether Value was derived from a global deposit
	// rule, as opposed to being set explicitly on this Price.
	Calculated bool    `json:"calculated"`
	Value      float64 `json:"value"`
}

// Price is the cost of booking a single [ElementOption] on a specific
// Departure, in one operator currency. Exactly one Price exists per
// operator currency for every departure/option pairing; see
// [DepartureElementOption.Prices] to discover their IDs.
type Price struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Currency is the ISO 4217 currency code this Price is quoted in.
	Currency string `json:"currency"`
	// OnSale reports whether this Price is currently available for sale.
	// Configurations that are not possible or disallowed are disabled.
	OnSale bool `json:"on_sale"`

	// Initialized reports whether this Price has been given a value yet.
	//
	// Note: openapi.yml's Price schema documents "initialized" with a
	// description but no explicit `type` (a spec omission, not an absent
	// field — it's still listed under `properties`). bool is inferred from
	// its description text ("Determines if the price has been initialized
	// with a value").
	Initialized bool    `json:"initialized"`
	Value       float64 `json:"value"`

	// Deposit is nil if no deposit rule applies to this Price.
	Deposit *PriceDeposit `json:"deposit,omitempty"`
}

// UpdatePriceParams are the properties that may change on an existing
// Price. All fields are optional; a nil field is left unchanged.
//
// Deposit is `nullable: true` in openapi.yml: use [NullValue] to set an
// explicit deposit, [Null] to clear it (reverting to the operator's default
// calculation), or leave it nil to leave the deposit unchanged.
//
// spec: updatePrice
type UpdatePriceParams struct {
	Value   *float64           `json:"value,omitempty"`
	Deposit *Nullable[float64] `json:"deposit,omitempty"`
	OnSale  *bool              `json:"on_sale,omitempty"`
}

// PricesResource provides access to the Prices API.
//
// Obtain one via [Client.Prices]. There is no List/Create/Delete: Prices
// are generated automatically, one per operator currency, alongside an
// [ElementOption] — discover their IDs via [DepartureElementOption.Prices].
type PricesResource struct {
	client *Client
}

// Prices returns a resource handle for the Prices API.
func (c *Client) Prices() *PricesResource {
	return &PricesResource{client: c}
}

// Get returns a single Price by ID.
//
// spec: getPrice
func (r *PricesResource) Get(ctx context.Context, id string) (*Price, error) {
	var p Price
	if err := r.client.do(ctx, http.MethodGet, "/holidays/prices/"+url.PathEscape(id), nil, nil, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Update partially updates an existing Price's value, deposit or sale
// state.
//
// spec: updatePrice
func (r *PricesResource) Update(ctx context.Context, id string, params UpdatePriceParams) (*Price, error) {
	var p Price
	if err := r.client.do(ctx, http.MethodPatch, "/holidays/prices/"+url.PathEscape(id), nil, params, &p); err != nil {
		return nil, err
	}
	return &p, nil
}
