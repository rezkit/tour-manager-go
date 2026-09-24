package tourmanager

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestFieldsResource_List(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/fields/holiday" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/fields/list.json")
	}))

	groups, err := client.Fields().List(context.Background(), EntityTypeHoliday)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}

	details := groups[0]
	if details.Label != "Details" || len(details.Fields) != 3 {
		t.Fatalf("unexpected group 0: %+v", details)
	}

	text, ok := details.Fields[0].(TextFieldDefinition)
	if !ok {
		t.Fatalf("field 0 type = %T, want TextFieldDefinition", details.Fields[0])
	}
	if text.Name != "trip_notes" || text.Type != TextFieldTypePlain {
		t.Errorf("unexpected text field: %+v", text)
	}
	if text.MinLength == nil || *text.MinLength != 5 {
		t.Errorf("MinLength = %v, want 5", text.MinLength)
	}

	date, ok := details.Fields[1].(DateFieldDefinition)
	if !ok {
		t.Fatalf("field 1 type = %T, want DateFieldDefinition", details.Fields[1])
	}
	if !date.Required {
		t.Errorf("date.Required = false, want true")
	}

	number, ok := details.Fields[2].(NumberFieldDefinition)
	if !ok {
		t.Fatalf("field 2 type = %T, want NumberFieldDefinition", details.Fields[2])
	}
	if number.MinimumValue == nil || *number.MinimumValue != 1 {
		t.Errorf("MinimumValue = %v, want 1", number.MinimumValue)
	}

	prefs := groups[1]
	if len(prefs.Fields) != 2 {
		t.Fatalf("unexpected group 1: %+v", prefs)
	}
	if _, ok := prefs.Fields[0].(BooleanFieldDefinition); !ok {
		t.Fatalf("field 0 type = %T, want BooleanFieldDefinition", prefs.Fields[0])
	}
	sel, ok := prefs.Fields[1].(SelectionFieldDefinition)
	if !ok {
		t.Fatalf("field 1 type = %T, want SelectionFieldDefinition", prefs.Fields[1])
	}
	if sel.GroupID != "01gpkgcy6t0m84czh8gy4kgr2" || sel.Type != SelectionFieldTypeMultiple {
		t.Errorf("unexpected selection field: %+v", sel)
	}
	if len(sel.Values) != 3 {
		t.Errorf("unexpected selection values: %v", sel.Values)
	}
}

func TestFieldsResource_Create_text(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/fields/holiday" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusCreated, "testdata/fields/create_text.json")
	}))

	fd, err := client.Fields().Create(context.Background(), EntityTypeHoliday, CreateTextFieldParams{
		Name:     "trip_notes",
		GroupID:  "01gpkgcy6t0m84czh8gy4kgr1",
		Label:    "Trip Notes",
		Required: false,
		Type:     TextFieldTypePlain,
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	text, ok := fd.(TextFieldDefinition)
	if !ok {
		t.Fatalf("type = %T, want TextFieldDefinition", fd)
	}
	if text.Name != "trip_notes" {
		t.Errorf("unexpected field: %+v", text)
	}
	if body := string(gotBody); !strings.Contains(body, `"type":"text"`) {
		t.Errorf("unexpected request body: %s", body)
	}
}

func TestFieldsResource_Create_booleanInjectsDiscriminator(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusCreated, "testdata/fields/create_boolean.json")
	}))

	fd, err := client.Fields().Create(context.Background(), EntityTypeHoliday, CreateBooleanFieldParams{
		Name:    "wheelchair_access",
		GroupID: "01gpkgcy6t0m84czh8gy4kgr2",
		Label:   "Wheelchair Access",
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if _, ok := fd.(BooleanFieldDefinition); !ok {
		t.Fatalf("type = %T, want BooleanFieldDefinition", fd)
	}
	if body := string(gotBody); !strings.Contains(body, `"type":"boolean"`) {
		t.Errorf("expected synthesized type discriminator in body: %s", body)
	}
}

func TestFieldsResource_Create_validationError(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusUnprocessableEntity, "testdata/fields/validation_error.json")
	}))

	_, err := client.Fields().Create(context.Background(), EntityTypeHoliday, CreateTextFieldParams{})
	var verr *ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("expected *ValidationError, got %v (%T)", err, err)
	}
}

func TestFieldsResource_Update_text(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch || r.URL.Path != "/fields/holiday/trip_notes" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/fields/update_text.json")
	}))

	fd, err := client.Fields().Update(context.Background(), EntityTypeHoliday, "trip_notes", UpdateTextFieldParams{
		MinLength: ptr(Null[int]()),
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	text, ok := fd.(TextFieldDefinition)
	if !ok {
		t.Fatalf("type = %T, want TextFieldDefinition", fd)
	}
	if text.Label != "Trip Notes (updated)" {
		t.Errorf("unexpected field: %+v", text)
	}
	if body := string(gotBody); !strings.Contains(body, `"minLength":null`) {
		t.Errorf("expected explicit null minLength: %s", body)
	}
}

func TestFieldsResource_Update_number(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/fields/holiday/party_size" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/fields/update_number.json")
	}))

	fd, err := client.Fields().Update(context.Background(), EntityTypeHoliday, "party_size", UpdateNumberFieldParams{
		Precision: ptr(NullValue(0)),
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	if _, ok := fd.(NumberFieldDefinition); !ok {
		t.Fatalf("type = %T, want NumberFieldDefinition", fd)
	}
}

func TestFieldsResource_CreateGroup(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/fields/holiday/groups" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusCreated, "testdata/fields/group.json")
	}))

	g, err := client.Fields().CreateGroup(context.Background(), EntityTypeHoliday, CreateFieldGroupParams{
		Name: "Details",
	})
	if err != nil {
		t.Fatalf("CreateGroup: %v", err)
	}
	if g.Name != "Details" {
		t.Fatalf("unexpected group: %+v", g)
	}
}
