package tourmanager

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestElementsResource_List(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/holidays/versions/01gpkgcy6t0m84czh8gy4kver1/elements" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/elements/list.json")
	}))

	page, err := client.Elements().List(context.Background(), "01gpkgcy6t0m84czh8gy4kver1", ListElementsOptions{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(page.Data) != 2 || page.Data[0].Name != "Hotel Room" {
		t.Fatalf("unexpected elements: %+v", page.Data)
	}

	alloc, ok := page.Data[0].DefaultInventory.(AllocationInventory)
	if !ok {
		t.Fatalf("element 0: DefaultInventory type = %T, want AllocationInventory", page.Data[0].DefaultInventory)
	}
	if alloc.Capacity != 40 {
		t.Errorf("AllocationInventory.Capacity = %d, want 40", alloc.Capacity)
	}
	if page.Data[0].BalanceDue == nil || *page.Data[0].BalanceDue != 56 {
		t.Errorf("BalanceDue = %v, want 56", page.Data[0].BalanceDue)
	}

	if _, ok := page.Data[1].DefaultInventory.(FreeSellInventory); !ok {
		t.Fatalf("element 1: DefaultInventory type = %T, want FreeSellInventory", page.Data[1].DefaultInventory)
	}
	if page.Data[1].BalanceDue != nil {
		t.Errorf("element 1: BalanceDue = %v, want nil", page.Data[1].BalanceDue)
	}
}

func TestElementsResource_List_queryParams(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if got := q.Get("name"); got != "Hotel" {
			t.Errorf("name = %q, want %q", got, "Hotel")
		}
		if got := q.Get("sort"); got != "name" {
			t.Errorf("sort = %q, want %q", got, "name")
		}
		if got := q.Get("published"); got != "1" {
			t.Errorf("published = %q, want %q", got, "1")
		}
		serveFile(t, w, http.StatusOK, "testdata/elements/list.json")
	}))

	published := true
	_, err := client.Elements().List(context.Background(), "ver1", ListElementsOptions{
		Name:      "Hotel",
		Sort:      ElementSortName,
		Published: &published,
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
}

func TestElementsResource_Get(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := "/holidays/versions/01gpkgcy6t0m84czh8gy4kver1/elements/01gpkgcy6t0m84czh8gy4kel1"
		if r.URL.Path != wantPath {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/elements/get.json")
	}))

	e, err := client.Elements().Get(context.Background(), "01gpkgcy6t0m84czh8gy4kver1", "01gpkgcy6t0m84czh8gy4kel1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if e.Name != "Hotel Room" {
		t.Fatalf("unexpected element: %+v", e)
	}
	if len(e.Options) != 1 || e.Options[0].Name != "Double Room" {
		t.Fatalf("unexpected options: %+v", e.Options)
	}
}

func TestElementsResource_Get_notFound(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusNotFound, "testdata/elements/not_found.json")
	}))

	_, err := client.Elements().Get(context.Background(), "ver1", "missing")
	var nferr *NotFoundError
	if !errors.As(err, &nferr) {
		t.Fatalf("expected *NotFoundError, got %v (%T)", err, err)
	}
}

func TestElementsResource_Create(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/holidays/versions/ver1/elements" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusCreated, "testdata/elements/get.json")
	}))

	_, err := client.Elements().Create(context.Background(), "ver1", CreateElementParams{
		Name:             "Hotel Room",
		CategoryID:       "cat1",
		DefaultInventory: AllocationInventory{Capacity: 40},
		BalanceDue:       ptr(NullValue(56)),
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	body := string(gotBody)
	if !strings.Contains(body, `"balance_due":56`) || !strings.Contains(body, `"type":"allocation"`) {
		t.Errorf("unexpected request body: %s", body)
	}
}

func TestElementsResource_Update_clearsBalanceDue(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/elements/get.json")
	}))

	_, err := client.Elements().Update(context.Background(), "ver1", "el1", UpdateElementParams{
		BalanceDue: ptr(Null[int]()),
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	body := string(gotBody)
	if !strings.Contains(body, `"balance_due":null`) {
		t.Errorf("expected explicit null balance_due: %s", body)
	}
	if strings.Contains(body, `"name"`) {
		t.Errorf("expected name to be omitted (unset): %s", body)
	}
}

func TestElementsResource_Delete(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := "/holidays/versions/ver1/elements/el1"
		if r.Method != http.MethodDelete || r.URL.Path != wantPath {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	if err := client.Elements().Delete(context.Background(), "ver1", "el1"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
}

func TestElementsResource_Categories_delegatesToAttachmentHandle(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/element/el1/categories" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/categories/list.json")
	}))

	page, err := client.Elements().Categories("el1").List(context.Background(), ListCategoriesOptions{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(page.Data) != 2 {
		t.Fatalf("unexpected categories: %+v", page.Data)
	}
}
