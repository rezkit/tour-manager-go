package tourmanager

import (
	"encoding/json"
	"fmt"
)

// Inventory is implemented by [AllocationInventory], [OnRequestInventory]
// and [FreeSellInventory] — RezKit's three inventory strategies for a
// Departure. Use a type switch to access variant-specific fields:
//
//	switch inv := departure.Inventory.(type) {
//	case tourmanager.AllocationInventory:
//		fmt.Println(inv.Capacity)
//	case tourmanager.OnRequestInventory:
//		fmt.Println(inv.Errata)
//	case tourmanager.FreeSellInventory:
//		// no additional fields
//	}
//
// openapi.yml declares Inventory as a oneOf with an explicit discriminator
// (propertyName: type); the variants below are dispatched on that field.
type Inventory interface {
	inventoryType() string
}

// AllocationInventory is inventory based on a fixed allocation: the total
// number of units available is given by Capacity.
type AllocationInventory struct {
	// Capacity is the total number of units in inventory.
	Capacity int `json:"capacity"`
}

func (AllocationInventory) inventoryType() string { return "allocation" }

// MarshalJSON includes the "type" discriminator the API expects.
func (v AllocationInventory) MarshalJSON() ([]byte, error) {
	type alias AllocationInventory
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "allocation", alias: alias(v)})
}

// OnRequestInventory is inventory that is confirmed manually rather than
// held automatically. Errata describes the confirmation process.
type OnRequestInventory struct {
	Errata string `json:"errata"`
}

func (OnRequestInventory) inventoryType() string { return "on_request" }

// MarshalJSON includes the "type" discriminator the API expects.
func (v OnRequestInventory) MarshalJSON() ([]byte, error) {
	type alias OnRequestInventory
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "on_request", alias: alias(v)})
}

// FreeSellInventory is unlimited, always-available inventory.
type FreeSellInventory struct{}

func (FreeSellInventory) inventoryType() string { return "free_sell" }

// MarshalJSON includes the "type" discriminator the API expects.
func (v FreeSellInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
	}{Type: "free_sell"})
}

// unmarshalInventory decodes raw into the Inventory variant identified by
// its "type" discriminator field.
func unmarshalInventory(raw []byte) (Inventory, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	var disc struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &disc); err != nil {
		return nil, fmt.Errorf("tourmanager: inventory: %w", err)
	}
	switch disc.Type {
	case "allocation":
		var v AllocationInventory
		return v, json.Unmarshal(raw, &v)
	case "on_request":
		var v OnRequestInventory
		return v, json.Unmarshal(raw, &v)
	case "free_sell":
		var v FreeSellInventory
		return v, json.Unmarshal(raw, &v)
	default:
		return nil, fmt.Errorf("tourmanager: unknown inventory type %q", disc.Type)
	}
}
