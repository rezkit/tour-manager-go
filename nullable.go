package tourmanager

import "encoding/json"

// Nullable represents a field on a write request that can be left unset
// (omitted from the request body entirely), explicitly cleared (sent as
// JSON null), or set to a value. It exists for the handful of fields the
// API documents as `nullable: true` on an update body, where a plain
// pointer can't distinguish "don't touch this field" from "clear it".
//
// The zero value of Nullable[T] is unset. Use [NullValue] to set a value,
// or [Null] to explicitly clear the field.
type Nullable[T any] struct {
	value T
	valid bool // true if value should be sent, as opposed to JSON null
	set   bool // true if this field should be sent at all
}

// NullValue returns a Nullable[T] that sends v as the field's value.
func NullValue[T any](v T) Nullable[T] {
	return Nullable[T]{value: v, valid: true, set: true}
}

// Null returns a Nullable[T] that explicitly clears the field (sends JSON
// null), as opposed to the zero value, which omits the field entirely.
func Null[T any]() Nullable[T] {
	return Nullable[T]{set: true}
}

// MarshalJSON implements json.Marshaler. An unset Nullable marshals to
// JSON null; callers should only include a *Nullable[T] field in a struct
// with `json:"...,omitempty"` so an unset field is omitted from the
// request entirely rather than sent as null.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.set || !n.valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}
