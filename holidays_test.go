package tourmanager

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func serveFile(t *testing.T, w http.ResponseWriter, status int, path string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read fixture %s: %v", path, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func TestHolidaysResource_List(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/holidays" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-api-key" {
			t.Errorf("Authorization header = %q", got)
		}
		serveFile(t, w, http.StatusOK, "testdata/holidays/list_page1.json")
	}))

	page, err := client.Holidays().List(context.Background(), ListHolidaysOptions{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page.Total != 3 || page.CurrentPage != 1 || page.LastPage != 2 {
		t.Fatalf("unexpected page metadata: %+v", page)
	}
	if len(page.Data) != 2 || page.Data[0].Code != "ALP01" {
		t.Fatalf("unexpected page data: %+v", page.Data)
	}
}

func TestHolidaysResource_List_queryParams(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		wantPage, wantLimit, wantSort, wantOrder := "2", "50", "code", "asc"
		if got := q.Get("page"); got != wantPage {
			t.Errorf("page = %q, want %q", got, wantPage)
		}
		if got := q.Get("limit"); got != wantLimit {
			t.Errorf("limit = %q, want %q", got, wantLimit)
		}
		if got := q.Get("sort"); got != wantSort {
			t.Errorf("sort = %q, want %q", got, wantSort)
		}
		if got := q.Get("order"); got != wantOrder {
			t.Errorf("order = %q, want %q", got, wantOrder)
		}
		if got := q.Get("published"); got != "1" {
			t.Errorf("published = %q, want %q", got, "1")
		}
		serveFile(t, w, http.StatusOK, "testdata/holidays/list_page1.json")
	}))

	published := true
	_, err := client.Holidays().List(context.Background(), ListHolidaysOptions{
		ListOptions: ListOptions{Page: 2, Limit: 50},
		Sort:        HolidaySortCode,
		Order:       SortAsc,
		Published:   &published,
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
}

func TestHolidaysResource_All(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("page") {
		case "", "1":
			serveFile(t, w, http.StatusOK, "testdata/holidays/list_page1.json")
		case "2":
			serveFile(t, w, http.StatusOK, "testdata/holidays/list_page2.json")
		default:
			t.Fatalf("unexpected page: %s", r.URL.Query().Get("page"))
		}
	}))

	var codes []string
	for holiday, err := range client.Holidays().All(context.Background(), ListHolidaysOptions{}) {
		if err != nil {
			t.Fatalf("All: %v", err)
		}
		codes = append(codes, holiday.Code)
	}

	want := []string{"ALP01", "CST01", "DST01"}
	if len(codes) != len(want) {
		t.Fatalf("got %v, want %v", codes, want)
	}
	for i := range want {
		if codes[i] != want[i] {
			t.Fatalf("got %v, want %v", codes, want)
		}
	}
}

func TestHolidaysResource_All_stopsOnError(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("page") {
		case "", "1":
			serveFile(t, w, http.StatusOK, "testdata/holidays/list_page1.json")
		default:
			serveFile(t, w, http.StatusNotFound, "testdata/holidays/not_found.json")
		}
	}))

	var codes []string
	var gotErr error
	for holiday, err := range client.Holidays().All(context.Background(), ListHolidaysOptions{}) {
		if err != nil {
			gotErr = err
			break
		}
		codes = append(codes, holiday.Code)
	}

	if len(codes) != 2 {
		t.Fatalf("got %d items before error, want 2: %v", len(codes), codes)
	}
	var nferr *NotFoundError
	if !errors.As(gotErr, &nferr) {
		t.Fatalf("expected *NotFoundError, got %v (%T)", gotErr, gotErr)
	}
}

func TestHolidaysResource_Get(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/holidays/01gpkgcy6t0m84czh8gy4kjat1" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/holidays/get.json")
	}))

	h, err := client.Holidays().Get(context.Background(), "01gpkgcy6t0m84czh8gy4kjat1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if h.Name != "Alpine Explorer" {
		t.Fatalf("unexpected holiday: %+v", h)
	}
}

func TestHolidaysResource_Get_customFields(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusOK, "testdata/holidays/get_with_fields.json")
	}))

	h, err := client.Holidays().Get(context.Background(), "01gpkgcy6t0m84czh8gy4kjat1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if len(h.Fields) != 2 {
		t.Fatalf("expected 2 custom fields, got %d: %+v", len(h.Fields), h.Fields)
	}

	notes, ok := h.Fields["trip_notes"].(TextFieldValue)
	if !ok || notes.Value != "Bring hiking boots" {
		t.Errorf("unexpected trip_notes field: %+v (ok=%v)", h.Fields["trip_notes"], ok)
	}
	access, ok := h.Fields["wheelchair_access"].(BooleanFieldValue)
	if !ok || access.Value != false {
		t.Errorf("unexpected wheelchair_access field: %+v (ok=%v)", h.Fields["wheelchair_access"], ok)
	}
}

func TestHolidaysResource_Get_notFound(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusNotFound, "testdata/holidays/not_found.json")
	}))

	_, err := client.Holidays().Get(context.Background(), "missing")
	var nferr *NotFoundError
	if !errors.As(err, &nferr) {
		t.Fatalf("expected *NotFoundError, got %v (%T)", err, err)
	}
	if nferr.StatusCode() != http.StatusNotFound {
		t.Fatalf("StatusCode = %d, want 404", nferr.StatusCode())
	}
	if nferr.Message != "Resource not found" {
		t.Fatalf("Message = %q", nferr.Message)
	}
}

func TestHolidaysResource_Create_validationError(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/holidays" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusUnprocessableEntity, "testdata/holidays/validation_error.json")
	}))

	_, err := client.Holidays().Create(context.Background(), CreateHolidayParams{Name: "x", Code: "y"})
	var verr *ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("expected *ValidationError, got %v (%T)", err, err)
	}
	if len(verr.Errors["name"]) == 0 || len(verr.Errors["code"]) == 0 {
		t.Fatalf("unexpected field errors: %+v", verr.Errors)
	}
}

func TestHolidaysResource_Update_nullableFields(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/holidays/get.json")
	}))

	_, err := client.Holidays().Update(context.Background(), "id", UpdateHolidayParams{
		Introduction: ptr(Null[string]()),
		Description:  ptr(NullValue("new description")),
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	body := string(gotBody)
	if !strings.Contains(body, `"introduction":null`) {
		t.Errorf("expected explicit null introduction in body: %s", body)
	}
	if !strings.Contains(body, `"description":"new description"`) {
		t.Errorf("expected set description in body: %s", body)
	}
	if strings.Contains(body, `"name"`) {
		t.Errorf("expected name to be omitted (unset): %s", body)
	}
}

func ptr[T any](v T) *T { return &v }
