package tourmanager

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestDeparturesResource_List_inventoryVariants(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/holidays/departures" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/departures/list.json")
	}))

	page, err := client.Departures().List(context.Background(), ListDeparturesOptions{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(page.Data) != 3 {
		t.Fatalf("expected 3 departures, got %d", len(page.Data))
	}

	alloc, ok := page.Data[0].Inventory.(AllocationInventory)
	if !ok {
		t.Fatalf("departure 0: Inventory type = %T, want AllocationInventory", page.Data[0].Inventory)
	}
	if alloc.Capacity != 20 {
		t.Errorf("AllocationInventory.Capacity = %d, want 20", alloc.Capacity)
	}

	onReq, ok := page.Data[1].Inventory.(OnRequestInventory)
	if !ok {
		t.Fatalf("departure 1: Inventory type = %T, want OnRequestInventory", page.Data[1].Inventory)
	}
	if onReq.Errata != "Contact us to confirm availability" {
		t.Errorf("OnRequestInventory.Errata = %q", onReq.Errata)
	}

	if _, ok := page.Data[2].Inventory.(FreeSellInventory); !ok {
		t.Fatalf("departure 2: Inventory type = %T, want FreeSellInventory", page.Data[2].Inventory)
	}
}

func TestDeparturesResource_List_queryParams(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if got := q.Get("holiday"); got != "hol123" {
			t.Errorf("holiday = %q", got)
		}
		if got := q.Get("published"); got != "1" {
			t.Errorf("published = %q", got)
		}
		serveFile(t, w, http.StatusOK, "testdata/departures/list.json")
	}))

	published := true
	_, err := client.Departures().List(context.Background(), ListDeparturesOptions{
		HolidayID: "hol123",
		Published: &published,
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
}

func TestDeparturesResource_Get(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusOK, "testdata/departures/get.json")
	}))

	d, err := client.Departures().Get(context.Background(), "01gpkgcy6t0m84czh8gy4kdep1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if _, ok := d.Inventory.(AllocationInventory); !ok {
		t.Fatalf("Inventory type = %T", d.Inventory)
	}
}

func TestCreateDepartureParams_marshalsInventoryDiscriminator(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusCreated, "testdata/departures/get.json")
	}))

	_, err := client.Departures().Create(context.Background(), CreateDepartureParams{
		Inventory: AllocationInventory{Capacity: 5},
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(gotBody, &decoded); err != nil {
		t.Fatalf("decode request body: %v", err)
	}
	inv, ok := decoded["inventory"].(map[string]any)
	if !ok {
		t.Fatalf("inventory field missing or wrong shape: %v", decoded)
	}
	if inv["type"] != "allocation" || inv["capacity"] != float64(5) {
		t.Fatalf("unexpected inventory payload: %v", inv)
	}
}
