package tourmanager

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCustomFieldsData_UnmarshalJSON_allVariants(t *testing.T) {
	const body = `{
		"notes": {"type": "text", "value": "Bring hiking boots"},
		"travel_date": {"type": "date", "value": "2024-06-01"},
		"party_size": {"type": "number", "value": 4},
		"wheelchair_access": {"type": "boolean", "value": true},
		"room_pref": {"type": "multi_select", "value": ["twin", "double"]}
	}`

	var data CustomFieldsData
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(data) != 5 {
		t.Fatalf("expected 5 fields, got %d", len(data))
	}

	text, ok := data["notes"].(TextFieldValue)
	if !ok || text.Value != "Bring hiking boots" || text.Type != TextFieldTypePlain {
		t.Errorf("unexpected notes field: %+v (ok=%v)", data["notes"], ok)
	}

	date, ok := data["travel_date"].(DateFieldValue)
	if !ok || date.Value != "2024-06-01" {
		t.Errorf("unexpected travel_date field: %+v (ok=%v)", data["travel_date"], ok)
	}

	number, ok := data["party_size"].(NumberFieldValue)
	if !ok || number.Value != 4 {
		t.Errorf("unexpected party_size field: %+v (ok=%v)", data["party_size"], ok)
	}

	boolean, ok := data["wheelchair_access"].(BooleanFieldValue)
	if !ok || !boolean.Value {
		t.Errorf("unexpected wheelchair_access field: %+v (ok=%v)", data["wheelchair_access"], ok)
	}

	sel, ok := data["room_pref"].(SelectFieldValue)
	if !ok || sel.Type != SelectFieldValueTypeMultiple || len(sel.Values) != 2 {
		t.Errorf("unexpected room_pref field: %+v (ok=%v)", data["room_pref"], ok)
	}
}

func TestCustomFieldsData_UnmarshalJSON_unknownType(t *testing.T) {
	var data CustomFieldsData
	err := json.Unmarshal([]byte(`{"x": {"type": "unknown", "value": 1}}`), &data)
	if err == nil {
		t.Fatal("expected an error for an unknown field value type")
	}
}

func TestCustomFieldsData_MarshalJSON_synthesizedTypes(t *testing.T) {
	data := CustomFieldsData{
		"travel_date":       DateFieldValue{Value: "2024-06-01"},
		"party_size":        NumberFieldValue{Value: 4},
		"wheelchair_access": BooleanFieldValue{Value: true},
	}

	out, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	body := string(out)
	for _, want := range []string{`"type":"date"`, `"type":"number"`, `"type":"boolean"`} {
		if !strings.Contains(body, want) {
			t.Errorf("expected %s in marshaled body: %s", want, body)
		}
	}

	// Round-trip: decoding what we just encoded should reproduce equivalent
	// values.
	var decoded CustomFieldsData
	if err := json.Unmarshal(out, &decoded); err != nil {
		t.Fatalf("round-trip Unmarshal: %v", err)
	}
	if decoded["party_size"].(NumberFieldValue).Value != 4 {
		t.Errorf("round-trip party_size = %+v", decoded["party_size"])
	}
}
