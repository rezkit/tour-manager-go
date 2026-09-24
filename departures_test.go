package tourmanager

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
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

func TestDeparturesResource_Get_elementsTree(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusOK, "testdata/departures/get_with_elements.json")
	}))

	d, err := client.Departures().Get(context.Background(), "01gpkgcy6t0m84czh8gy4kdep1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if len(d.Elements) != 1 {
		t.Fatalf("expected 1 departure element, got %d", len(d.Elements))
	}

	de := d.Elements[0]
	if de.ID != "01gpkgcy6t0m84czh8gy4kdel1" {
		t.Errorf("DepartureElement.ID = %q", de.ID)
	}
	onReq, ok := de.Inventory.(OnRequestInventory)
	if !ok {
		t.Fatalf("DepartureElement.Inventory type = %T, want OnRequestInventory", de.Inventory)
	}
	if onReq.Errata != "Confirm with supplier" {
		t.Errorf("OnRequestInventory.Errata = %q", onReq.Errata)
	}
	if de.Element.ID != "01gpkgcy6t0m84czh8gy4kel1" || de.Element.Name != "Hotel Room" {
		t.Errorf("unexpected element summary: %+v", de.Element)
	}
	if de.BalanceDue == nil || !de.BalanceDue.Calculated {
		t.Fatalf("unexpected balance due: %+v", de.BalanceDue)
	}

	if len(de.Options) != 1 {
		t.Fatalf("expected 1 departure element option, got %d", len(de.Options))
	}
	opt := de.Options[0]
	if opt.ID != "01gpkgcy6t0m84czh8gy4kopt1" || opt.Name != "Double Room" {
		t.Fatalf("unexpected departure element option: %+v", opt)
	}
	if len(opt.Prices) != 2 {
		t.Fatalf("expected 2 prices, got %d", len(opt.Prices))
	}
	if opt.Prices[0].ID != "01gpkgcy6t0m84czh8gy4kprc1" || opt.Prices[0].Currency != "USD" {
		t.Errorf("unexpected price 0: %+v", opt.Prices[0])
	}
	if opt.Prices[1].Currency != "GBP" || opt.Prices[1].Deposit != nil {
		t.Errorf("unexpected price 1: %+v", opt.Prices[1])
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
		VersionID: "01gpkgcy6t0m84czh8gy4kver1",
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
	if decoded["version_id"] != "01gpkgcy6t0m84czh8gy4kver1" {
		t.Fatalf("expected version_id in request body: %v", decoded)
	}
}

func TestCreateDepartureParams_publishedAndFields(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusCreated, "testdata/departures/get.json")
	}))

	published := true
	_, err := client.Departures().Create(context.Background(), CreateDepartureParams{
		VersionID: "01gpkgcy6t0m84czh8gy4kver1",
		Inventory: AllocationInventory{Capacity: 5},
		Published: &published,
		Fields: CustomFieldsData{
			"notes": TextFieldValue{Type: TextFieldTypePlain, Value: "Group booking"},
		},
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	body := string(gotBody)
	if !strings.Contains(body, `"published":true`) {
		t.Errorf("expected published in request body: %s", body)
	}
	if !strings.Contains(body, `"notes":{"type":"text","value":"Group booking"}`) {
		t.Errorf("expected fields in request body: %s", body)
	}
}

func TestUpdateDepartureParams_publishedAndFields(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusOK, "testdata/departures/get.json")
	}))

	published := false
	_, err := client.Departures().Update(context.Background(), "dep1", UpdateDepartureParams{
		Published: &published,
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	body := string(gotBody)
	if !strings.Contains(body, `"published":false`) {
		t.Errorf("expected published in request body: %s", body)
	}
	if strings.Contains(body, `"fields"`) || strings.Contains(body, `"inventory"`) {
		t.Errorf("expected fields/inventory to be omitted (unset): %s", body)
	}
}

func TestDepartureElements_Update_setsInventory(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := "/holidays/departures/dep1/elements/del1"
		if r.Method != http.MethodPatch || r.URL.Path != wantPath {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/departures/update_element.json")
	}))

	de, err := client.Departures().Elements("dep1").Update(context.Background(), "del1", UpdateDepartureElementParams{
		Inventory: AllocationInventory{Capacity: 15},
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	alloc, ok := de.Inventory.(AllocationInventory)
	if !ok || alloc.Capacity != 15 {
		t.Fatalf("unexpected inventory: %+v (ok=%v)", de.Inventory, ok)
	}
	if de.BalanceDue == nil || de.BalanceDue.Calculated {
		t.Fatalf("unexpected balance due: %+v", de.BalanceDue)
	}

	var decoded map[string]any
	if err := json.Unmarshal(gotBody, &decoded); err != nil {
		t.Fatalf("decode request body: %v", err)
	}
	inv, ok := decoded["inventory"].(map[string]any)
	if !ok || inv["type"] != "allocation" || inv["capacity"] != float64(15) {
		t.Fatalf("unexpected inventory payload: %v", decoded)
	}
	if _, present := decoded["balance_due"]; present {
		t.Errorf("expected balance_due to be omitted (unset): %v", decoded)
	}
}

func TestDepartureElements_Update_clearsBalanceDue(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/departures/update_element.json")
	}))

	_, err := client.Departures().Elements("dep1").Update(context.Background(), "del1", UpdateDepartureElementParams{
		BalanceDue: ptr(Null[time.Time]()),
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	body := string(gotBody)
	if !strings.Contains(body, `"balance_due":null`) {
		t.Errorf("expected explicit null balance_due: %s", body)
	}
	if strings.Contains(body, `"inventory"`) {
		t.Errorf("expected inventory to be omitted (unset): %s", body)
	}
}
