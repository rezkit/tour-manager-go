package tourmanager

import (
	"encoding/json"
	"fmt"
)

// CustomFieldsData holds the custom field values recorded against a
// specific entity (for example a Holiday, Location, Accommodation or
// Extra — every schema in openapi.yml that has a "fields" property of this
// type), keyed by field name. See [FieldsResource] for the field
// definitions this data is validated against.
type CustomFieldsData map[string]FieldValue

// UnmarshalJSON decodes CustomFieldsData, dispatching each entry to the
// concrete FieldValue type identified by that entry's own "type"
// discriminator.
func (d *CustomFieldsData) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	out := make(CustomFieldsData, len(raw))
	for name, v := range raw {
		fv, err := unmarshalFieldValue(v)
		if err != nil {
			return fmt.Errorf("tourmanager: custom field %q: %w", name, err)
		}
		out[name] = fv
	}
	*d = out
	return nil
}

// FieldValue is implemented by [TextFieldValue], [DateFieldValue],
// [NumberFieldValue], [BooleanFieldValue] and [SelectFieldValue] — the
// possible value types a custom field can hold, matching the
// [FieldDefinition] it was created against. Use a type switch to access
// variant-specific data.
type FieldValue interface {
	isFieldValue()
}

// TextFieldValue holds the value of a text-shaped custom field (any of
// TextFieldTypePlain, TextFieldTypeRichText, TextFieldTypeURL or
// TextFieldTypeColor).
type TextFieldValue struct {
	Type  TextFieldType `json:"type"`
	Value string        `json:"value"`
}

func (TextFieldValue) isFieldValue() {}

// DateFieldValue holds the value of a date-valued custom field. Value is
// kept as the raw date string the API sends (e.g. "2023-01-17") rather
// than decoded into time.Time, since this is a date-only format and no
// date-only time convention exists elsewhere in this client.
type DateFieldValue struct {
	Value string `json:"value"`
}

func (DateFieldValue) isFieldValue() {}

// MarshalJSON includes the "type" discriminator the API expects.
func (v DateFieldValue) MarshalJSON() ([]byte, error) {
	type alias DateFieldValue
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "date", alias: alias(v)})
}

// NumberFieldValue holds the value of a numeric custom field.
type NumberFieldValue struct {
	Value float64 `json:"value"`
}

func (NumberFieldValue) isFieldValue() {}

// MarshalJSON includes the "type" discriminator the API expects.
func (v NumberFieldValue) MarshalJSON() ([]byte, error) {
	type alias NumberFieldValue
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "number", alias: alias(v)})
}

// BooleanFieldValue holds the value of a boolean (checkbox) custom field.
type BooleanFieldValue struct {
	Value bool `json:"value"`
}

func (BooleanFieldValue) isFieldValue() {}

// MarshalJSON includes the "type" discriminator the API expects.
func (v BooleanFieldValue) MarshalJSON() ([]byte, error) {
	type alias BooleanFieldValue
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "boolean", alias: alias(v)})
}

// SelectFieldValueType distinguishes a single- from a multi-value
// [SelectFieldValue].
//
// Note: this is deliberately a different Go type from [SelectionFieldType]
// (used by [SelectionFieldDefinition]) — see that type's doc comment for
// why.
type SelectFieldValueType string

const (
	SelectFieldValueTypeSingle   SelectFieldValueType = "select"
	SelectFieldValueTypeMultiple SelectFieldValueType = "multi_select"
)

// SelectFieldValue holds the selected value(s) of a selection-shaped
// custom field. When Type is SelectFieldValueTypeSingle, Values contains
// zero or one entries; when SelectFieldValueTypeMultiple, it may contain
// more.
type SelectFieldValue struct {
	Type   SelectFieldValueType `json:"type"`
	Values []string             `json:"value"`
}

func (SelectFieldValue) isFieldValue() {}

// unmarshalFieldValue decodes raw into the FieldValue variant identified
// by its "type" field, matching against each variant's known set of enum
// values (not a single fixed value — TextFieldValue alone covers four).
func unmarshalFieldValue(raw []byte) (FieldValue, error) {
	var disc struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &disc); err != nil {
		return nil, fmt.Errorf("tourmanager: field value: %w", err)
	}
	switch disc.Type {
	case "text", "rich_text", "url", "color":
		var v TextFieldValue
		return v, json.Unmarshal(raw, &v)
	case "date":
		var v DateFieldValue
		return v, json.Unmarshal(raw, &v)
	case "number":
		var v NumberFieldValue
		return v, json.Unmarshal(raw, &v)
	case "boolean":
		var v BooleanFieldValue
		return v, json.Unmarshal(raw, &v)
	case "select", "multi_select":
		var v SelectFieldValue
		return v, json.Unmarshal(raw, &v)
	default:
		return nil, fmt.Errorf("tourmanager: unknown field value type %q", disc.Type)
	}
}
