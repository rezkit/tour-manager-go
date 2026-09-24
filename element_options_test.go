package tourmanager

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestElementOptions_Create(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/holidays/elements/el1/options" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusCreated, "testdata/element_options/get.json")
	}))

	opt, err := client.Elements().Options("el1").Create(context.Background(), CreateElementOptionParams{
		Name:       "Double Room",
		CategoryID: "cat3",
		Occupancy:  Occupancy{From: 2, To: 2},
		PriceUnit:  PriceUnitPerson,
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if opt.Name != "Double Room" || opt.PriceUnit != PriceUnitPerson {
		t.Fatalf("unexpected option: %+v", opt)
	}
	if opt.Constraints == nil || len(opt.Constraints.PassengerSex) != 2 {
		t.Fatalf("unexpected constraints: %+v", opt.Constraints)
	}
}

func TestElementOptions_Create_validationError(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusUnprocessableEntity, "testdata/element_options/validation_error.json")
	}))

	_, err := client.Elements().Options("el1").Create(context.Background(), CreateElementOptionParams{})
	var verr *ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("expected *ValidationError, got %v (%T)", err, err)
	}
	if len(verr.Errors["price_unit"]) == 0 {
		t.Fatalf("unexpected field errors: %+v", verr.Errors)
	}
}

func TestElementOptions_Get(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/holidays/elements/el1/options/opt1" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/element_options/get.json")
	}))

	opt, err := client.Elements().Options("el1").Get(context.Background(), "opt1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if opt.Occupancy != (Occupancy{From: 2, To: 2}) {
		t.Fatalf("unexpected occupancy: %+v", opt.Occupancy)
	}
}

func TestElementOptions_Update(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch || r.URL.Path != "/holidays/elements/el1/options/opt1" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/element_options/get.json")
	}))

	published := false
	_, err := client.Elements().Options("el1").Update(context.Background(), "opt1", UpdateElementOptionParams{
		Published: &published,
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
}

func TestElementOptions_Delete(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/holidays/elements/el1/options/opt1" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	if err := client.Elements().Options("el1").Delete(context.Background(), "opt1"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
}

func TestElementOptions_Restore(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/holidays/elements/el1/options/opt1/restore" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/element_options/get.json")
	}))

	opt, err := client.Elements().Options("el1").Restore(context.Background(), "opt1")
	if err != nil {
		t.Fatalf("Restore: %v", err)
	}
	if opt.Name != "Double Room" {
		t.Fatalf("unexpected option: %+v", opt)
	}
}
