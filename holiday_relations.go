package tourmanager

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// HolidayEdge represents a directed relationship between two Holidays.
type HolidayEdge struct {
	ID            string    `json:"id"`
	SourceID      string    `json:"source_id"`
	DestinationID string    `json:"destination_id"`
	CategoryID    string    `json:"category_id"`
	StartDay      int       `json:"start_day"`
	Published     bool      `json:"published"`
	Ordering      int       `json:"ordering"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// AddHolidayRelationParams are the properties for creating a relationship
// from one Holiday to another. DestinationID, CategoryID and Published are
// required.
type AddHolidayRelationParams struct {
	DestinationID string `json:"destination_id"`
	CategoryID    string `json:"category_id"`
	Published     bool   `json:"published"`

	StartDay *int `json:"start_day,omitempty"`
	// CreateTwin also creates the reverse relationship.
	CreateTwin *bool `json:"create_twin,omitempty"`
}

// UpdateHolidayRelationParams are the properties that may change on an
// existing holiday relation. All fields are optional.
type UpdateHolidayRelationParams struct {
	CategoryID *string          `json:"category_id,omitempty"`
	StartDay   *int             `json:"start_day,omitempty"`
	Published  *bool            `json:"published,omitempty"`
	Ordering   *OrderingCommand `json:"ordering,omitempty"`
}

// HolidayRelations manages relationships ("edges") between one specific
// Holiday and others.
//
// Obtain one via [HolidaysResource.Relations].
type HolidayRelations struct {
	client    *Client
	holidayID string
}

func (r *HolidayRelations) path(suffix string) string {
	return "/holidays/" + url.PathEscape(r.holidayID) + "/relations" + suffix
}

// List returns every Holiday related to this Holiday.
//
// The API does not paginate this endpoint — it returns a bare JSON array —
// so this returns []HolidayEdge rather than *Page[HolidayEdge].
//
// spec: listRelatedHolidays
func (r *HolidayRelations) List(ctx context.Context) ([]HolidayEdge, error) {
	var edges []HolidayEdge
	err := r.client.do(ctx, http.MethodGet, r.path(""), nil, nil, &edges)
	return edges, err
}

// Add creates a new relationship from this Holiday to another.
//
// spec: addRelatedHoliday
func (r *HolidayRelations) Add(ctx context.Context, params AddHolidayRelationParams) (*HolidayEdge, error) {
	var edge HolidayEdge
	if err := r.client.do(ctx, http.MethodPost, r.path(""), nil, params, &edge); err != nil {
		return nil, err
	}
	return &edge, nil
}

// Get returns a single relation by its edge ID.
//
// spec: getRelatedHoliday
func (r *HolidayRelations) Get(ctx context.Context, relationID string) (*HolidayEdge, error) {
	var edge HolidayEdge
	if err := r.client.do(ctx, http.MethodGet, r.path("/"+url.PathEscape(relationID)), nil, nil, &edge); err != nil {
		return nil, err
	}
	return &edge, nil
}

// Update changes properties of an existing relation.
//
// spec: updateRelatedHoliday
func (r *HolidayRelations) Update(ctx context.Context, relationID string, params UpdateHolidayRelationParams) (*HolidayEdge, error) {
	var edge HolidayEdge
	if err := r.client.do(ctx, http.MethodPatch, r.path("/"+url.PathEscape(relationID)), nil, params, &edge); err != nil {
		return nil, err
	}
	return &edge, nil
}

// Delete removes a relation between Holidays.
//
// spec: deleteRelatedHoliday
func (r *HolidayRelations) Delete(ctx context.Context, relationID string) error {
	return r.client.do(ctx, http.MethodDelete, r.path("/"+url.PathEscape(relationID)), nil, nil, nil)
}
