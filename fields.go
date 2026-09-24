package tourmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// TextFieldType is the specific input semantics of a text-shaped
// [FieldDefinition] or [FieldValue]. All four variants store a string
// value; Type only changes how it's presented/validated.
type TextFieldType string

const (
	TextFieldTypePlain    TextFieldType = "text"
	TextFieldTypeRichText TextFieldType = "rich_text"
	TextFieldTypeURL      TextFieldType = "url"
	TextFieldTypeColor    TextFieldType = "color"
)

// SelectionFieldType distinguishes a single- from a multi-value
// [SelectionFieldDefinition].
//
// Note: this is deliberately a different Go type from [SelectFieldValueType]
// (used by [SelectFieldValue]). openapi.yml documents the same concept with
// two different second enum values — "select_multi" here vs "multi_select"
// on the value side — which looks like a spec inconsistency, not two
// different things; each is kept exactly as documented rather than merged.
type SelectionFieldType string

const (
	SelectionFieldTypeSingle   SelectionFieldType = "select"
	SelectionFieldTypeMultiple SelectionFieldType = "select_multi"
)

// FieldDefinition is implemented by [TextFieldDefinition],
// [DateFieldDefinition], [NumberFieldDefinition], [BooleanFieldDefinition]
// and [SelectionFieldDefinition] — the configuration of one custom field
// belonging to an [EntityType]. Use a type switch to access variant-specific
// properties. See [CustomFieldsData] for the recorded values of fields
// defined this way.
type FieldDefinition interface {
	isFieldDefinition()
}

// TextFieldDefinition configures a text-shaped custom field.
type TextFieldDefinition struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	GroupID  string        `json:"group_id"`
	Rank     int           `json:"rank"`
	Label    string        `json:"label"`
	Required bool          `json:"required"`
	Type     TextFieldType `json:"type"`

	MinLength *int `json:"minLength,omitempty"`
	MaxLength *int `json:"maxLength,omitempty"`
	// Format is a regular expression the value must match. Only applicable
	// when Type is TextFieldTypePlain.
	Format string `json:"format,omitempty"`
}

func (TextFieldDefinition) isFieldDefinition() {}

// DateFieldDefinition configures a date-valued custom field.
type DateFieldDefinition struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	GroupID  string `json:"group_id"`
	Rank     int    `json:"rank"`
	Label    string `json:"label"`
	Required bool   `json:"required"`
}

func (DateFieldDefinition) isFieldDefinition() {}

// NumberFieldDefinition configures a numeric custom field.
type NumberFieldDefinition struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	GroupID   string `json:"group_id"`
	Rank      int    `json:"rank"`
	Label     string `json:"label"`
	Required  bool   `json:"required"`
	Precision int    `json:"precision"`

	// MinimumValue and MaximumValue are inclusive bounds, if set.
	MinimumValue *int `json:"minimumValue,omitempty"`
	MaximumValue *int `json:"maximumValue,omitempty"`
}

func (NumberFieldDefinition) isFieldDefinition() {}

// BooleanFieldDefinition configures a boolean (checkbox) custom field.
type BooleanFieldDefinition struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	GroupID string `json:"group_id"`
	Rank    int    `json:"rank"`
	Label   string `json:"label"`
}

func (BooleanFieldDefinition) isFieldDefinition() {}

// SelectionFieldDefinition configures a single- or multi-select custom
// field, constrained to Values.
type SelectionFieldDefinition struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Label string `json:"label"`

	// GroupID is sent under the JSON key "group" for this variant only —
	// every other FieldDefinition variant uses "group_id". Kept exactly as
	// openapi.yml documents it; likely a spec naming inconsistency, not a
	// different concept.
	GroupID string `json:"group"`

	Required bool               `json:"required"`
	Type     SelectionFieldType `json:"type"`
	Values   []string           `json:"values"`
}

func (SelectionFieldDefinition) isFieldDefinition() {}

// unmarshalFieldDefinition decodes raw into the FieldDefinition variant
// identified by its "type" field, matching against each variant's known set
// of enum values (not a single fixed value — TextFieldDefinition alone
// covers four).
func unmarshalFieldDefinition(raw []byte) (FieldDefinition, error) {
	var disc struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &disc); err != nil {
		return nil, fmt.Errorf("tourmanager: field definition: %w", err)
	}
	switch disc.Type {
	case "text", "rich_text", "url", "color":
		var v TextFieldDefinition
		return v, json.Unmarshal(raw, &v)
	case "date":
		var v DateFieldDefinition
		return v, json.Unmarshal(raw, &v)
	case "number":
		var v NumberFieldDefinition
		return v, json.Unmarshal(raw, &v)
	case "boolean":
		var v BooleanFieldDefinition
		return v, json.Unmarshal(raw, &v)
	case "select", "select_multi":
		var v SelectionFieldDefinition
		return v, json.Unmarshal(raw, &v)
	default:
		return nil, fmt.Errorf("tourmanager: unknown field definition type %q", disc.Type)
	}
}

// FieldGroup is a named grouping of custom fields, as returned by
// [FieldsResource.CreateGroup].
//
// Note: this is a different response shape from the group objects returned
// by [FieldsResource.List] (see [FieldGroupFields]) — openapi.yml documents
// them as two distinct, unnamed-vs-named schemas for a similar-sounding
// concept, not a spec bug to reconcile.
type FieldGroup struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Rank      int       `json:"rank"`
}

// FieldGroupFields is one field group and the FieldDefinitions it contains,
// as returned by [FieldsResource.List]. openapi.yml documents this as an
// untyped inline array item (id/label/rank/fields), distinct from the named
// FieldGroup schema returned by CreateGroup.
type FieldGroupFields struct {
	ID     string
	Label  string
	Rank   int
	Fields []FieldDefinition
}

// UnmarshalJSON decodes a FieldGroupFields, dispatching each entry in
// Fields to its concrete FieldDefinition type.
func (g *FieldGroupFields) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID     string            `json:"id"`
		Label  string            `json:"label"`
		Rank   int               `json:"rank"`
		Fields []json.RawMessage `json:"fields"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	fields := make([]FieldDefinition, len(aux.Fields))
	for i, raw := range aux.Fields {
		fd, err := unmarshalFieldDefinition(raw)
		if err != nil {
			return err
		}
		fields[i] = fd
	}

	g.ID, g.Label, g.Rank, g.Fields = aux.ID, aux.Label, aux.Rank, fields
	return nil
}

// CreateFieldParams is implemented by [CreateTextFieldParams],
// [CreateDateFieldParams], [CreateNumberFieldParams],
// [CreateBooleanFieldParams] and [CreateSelectionFieldParams].
//
// openapi.yml's createField request body is documented as the full
// FieldDefinition read schema, including "id" as required — reusing the
// read schema for the write body, which a client can't honor since the ID
// is server-generated. Per this client's usual convention, id is excluded
// from every variant here, matching every other Create<X>Params type.
type CreateFieldParams interface {
	isCreateFieldParams()
}

// CreateTextFieldParams creates a text-shaped custom field.
//
// spec: createField
type CreateTextFieldParams struct {
	Name     string        `json:"name"`
	GroupID  string        `json:"group_id"`
	Label    string        `json:"label"`
	Required bool          `json:"required"`
	Type     TextFieldType `json:"type"`

	Rank      *int    `json:"rank,omitempty"`
	MinLength *int    `json:"minLength,omitempty"`
	MaxLength *int    `json:"maxLength,omitempty"`
	Format    *string `json:"format,omitempty"`
}

func (CreateTextFieldParams) isCreateFieldParams() {}

// CreateDateFieldParams creates a date-valued custom field.
//
// spec: createField
type CreateDateFieldParams struct {
	Name     string `json:"name"`
	GroupID  string `json:"group_id"`
	Label    string `json:"label"`
	Required bool   `json:"required"`

	Rank *int `json:"rank,omitempty"`
}

func (CreateDateFieldParams) isCreateFieldParams() {}

// MarshalJSON includes the "type" discriminator the API expects.
func (p CreateDateFieldParams) MarshalJSON() ([]byte, error) {
	type alias CreateDateFieldParams
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "date", alias: alias(p)})
}

// CreateNumberFieldParams creates a numeric custom field.
//
// spec: createField
type CreateNumberFieldParams struct {
	Name      string `json:"name"`
	GroupID   string `json:"group_id"`
	Label     string `json:"label"`
	Required  bool   `json:"required"`
	Precision int    `json:"precision"`

	MinimumValue *int `json:"minimumValue,omitempty"`
	MaximumValue *int `json:"maximumValue,omitempty"`
	Rank         *int `json:"rank,omitempty"`
}

func (CreateNumberFieldParams) isCreateFieldParams() {}

// MarshalJSON includes the "type" discriminator the API expects.
func (p CreateNumberFieldParams) MarshalJSON() ([]byte, error) {
	type alias CreateNumberFieldParams
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "number", alias: alias(p)})
}

// CreateBooleanFieldParams creates a boolean (checkbox) custom field.
//
// spec: createField
type CreateBooleanFieldParams struct {
	Name    string `json:"name"`
	GroupID string `json:"group_id"`
	Label   string `json:"label"`

	Rank *int `json:"rank,omitempty"`
}

func (CreateBooleanFieldParams) isCreateFieldParams() {}

// MarshalJSON includes the "type" discriminator the API expects.
func (p CreateBooleanFieldParams) MarshalJSON() ([]byte, error) {
	type alias CreateBooleanFieldParams
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "boolean", alias: alias(p)})
}

// CreateSelectionFieldParams creates a single- or multi-select custom
// field, constrained to Values.
//
// spec: createField
type CreateSelectionFieldParams struct {
	Name string `json:"name"`
	// GroupID is sent under the JSON key "group"; see
	// SelectionFieldDefinition.GroupID.
	GroupID  string             `json:"group"`
	Label    string             `json:"label"`
	Required bool               `json:"required"`
	Type     SelectionFieldType `json:"type"`
	Values   []string           `json:"values"`

	Rank *int `json:"rank,omitempty"`
}

func (CreateSelectionFieldParams) isCreateFieldParams() {}

// UpdateFieldParams is implemented by [UpdateTextFieldParams] and
// [UpdateNumberFieldParams] — the only two field-update shapes
// openapi.yml documents. Date, Boolean and Selection fields have no
// documented update body at all (not even an untyped one) and so cannot be
// updated through this client yet; see README.md's "Known gaps".
type UpdateFieldParams interface {
	isUpdateFieldParams()
}

// UpdateTextFieldParams are the properties that may change on an existing
// text-shaped field. All fields are optional; a nil field is left
// unchanged. MinLength and MaxLength are `nullable: true` in openapi.yml:
// use [NullValue] to set a value, [Null] to explicitly clear it, or leave
// the pointer nil to leave it unchanged.
//
// spec: updateField
type UpdateTextFieldParams struct {
	GroupID *string `json:"group_id,omitempty"`
	Label   *string `json:"label,omitempty"`

	MinLength *Nullable[int] `json:"minLength,omitempty"`
	MaxLength *Nullable[int] `json:"maxLength,omitempty"`
}

func (UpdateTextFieldParams) isUpdateFieldParams() {}

// UpdateNumberFieldParams are the properties that may change on an
// existing numeric field. All fields are optional; a nil field is left
// unchanged. MinimumValue, MaximumValue and Precision are `nullable: true`
// in openapi.yml: use [NullValue]/[Null] as with [UpdateTextFieldParams].
//
// spec: updateField
type UpdateNumberFieldParams struct {
	GroupID *string `json:"group_id,omitempty"`
	Label   *string `json:"label,omitempty"`

	MinimumValue *Nullable[int] `json:"minimumValue,omitempty"`
	MaximumValue *Nullable[int] `json:"maximumValue,omitempty"`
	Precision    *Nullable[int] `json:"precision,omitempty"`
}

func (UpdateNumberFieldParams) isUpdateFieldParams() {}

// CreateFieldGroupParams are the properties for creating a new FieldGroup.
// Name is required.
//
// spec: createGroup
type CreateFieldGroupParams struct {
	Name string `json:"name"`
	Rank *int   `json:"rank,omitempty"`
}

// FieldsResource provides access to the custom Fields API: field
// definitions and field groups, both scoped by [EntityType] rather than by
// a specific item (all entities of a given type share the same field
// definitions). See [CustomFieldsData] for the per-item recorded values.
//
// Obtain one via [Client.Fields].
type FieldsResource struct {
	client *Client
}

// Fields returns a resource handle for the custom Fields API.
func (c *Client) Fields() *FieldsResource {
	return &FieldsResource{client: c}
}

// List returns every field group (and the FieldDefinitions it contains)
// configured for entityType. The API returns a bare JSON array with no
// pagination metadata, so this returns []FieldGroupFields rather than
// *Page[FieldGroupFields].
//
// spec: getFields
func (r *FieldsResource) List(ctx context.Context, entityType EntityType) ([]FieldGroupFields, error) {
	var groups []FieldGroupFields
	path := "/fields/" + url.PathEscape(string(entityType))
	err := r.client.do(ctx, http.MethodGet, path, nil, nil, &groups)
	return groups, err
}

// Create creates a new custom field on entityType.
//
// spec: createField
func (r *FieldsResource) Create(ctx context.Context, entityType EntityType, params CreateFieldParams) (FieldDefinition, error) {
	var raw json.RawMessage
	path := "/fields/" + url.PathEscape(string(entityType))
	if err := r.client.do(ctx, http.MethodPost, path, nil, params, &raw); err != nil {
		return nil, err
	}
	return unmarshalFieldDefinition(raw)
}

// Update partially updates an existing custom field, identified by its
// name.
//
// spec: updateField
func (r *FieldsResource) Update(ctx context.Context, entityType EntityType, name string, params UpdateFieldParams) (FieldDefinition, error) {
	var raw json.RawMessage
	path := "/fields/" + url.PathEscape(string(entityType)) + "/" + url.PathEscape(name)
	if err := r.client.do(ctx, http.MethodPatch, path, nil, params, &raw); err != nil {
		return nil, err
	}
	return unmarshalFieldDefinition(raw)
}

// CreateGroup creates a new FieldGroup on entityType.
//
// spec: createGroup
func (r *FieldsResource) CreateGroup(ctx context.Context, entityType EntityType, params CreateFieldGroupParams) (*FieldGroup, error) {
	var g FieldGroup
	path := "/fields/" + url.PathEscape(string(entityType)) + "/groups"
	if err := r.client.do(ctx, http.MethodPost, path, nil, params, &g); err != nil {
		return nil, err
	}
	return &g, nil
}
