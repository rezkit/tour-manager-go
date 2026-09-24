package tourmanager

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestHolidayVersions_List(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/holidays/hol1/versions" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/holiday_versions/list.json")
	}))

	page, err := client.Holidays().Versions("hol1").List(context.Background(), ListHolidayVersionsOptions{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(page.Data) != 2 || page.Data[0].Code != "HV2024" {
		t.Fatalf("unexpected versions: %+v", page.Data)
	}
	if page.Data[1].Published {
		t.Errorf("version 1: Published = true, want false")
	}
}

func TestHolidayVersions_List_queryParams(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if got := q.Get("sort"); got != "code" {
			t.Errorf("sort = %q, want %q", got, "code")
		}
		if got := q.Get("published"); got != "1" {
			t.Errorf("published = %q, want %q", got, "1")
		}
		serveFile(t, w, http.StatusOK, "testdata/holiday_versions/list.json")
	}))

	published := true
	_, err := client.Holidays().Versions("hol1").List(context.Background(), ListHolidayVersionsOptions{
		Sort:      HolidayVersionSortCode,
		Published: &published,
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
}

func TestHolidayVersions_All(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusOK, "testdata/holiday_versions/list.json")
	}))

	var codes []string
	for v, err := range client.Holidays().Versions("hol1").All(context.Background(), ListHolidayVersionsOptions{}) {
		if err != nil {
			t.Fatalf("All: %v", err)
		}
		codes = append(codes, v.Code)
	}
	if len(codes) != 2 {
		t.Fatalf("unexpected codes: %v", codes)
	}
}

func TestHolidayVersions_Get(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/holidays/hol1/versions/ver1" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/holiday_versions/get.json")
	}))

	v, err := client.Holidays().Versions("hol1").Get(context.Background(), "ver1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if v.HolidayID != "01gpkgcy6t0m84czh8gy4kjat1" || v.Name != "2024 Season" {
		t.Fatalf("unexpected version: %+v", v)
	}
}

func TestHolidayVersions_Get_notFound(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusNotFound, "testdata/holiday_versions/not_found.json")
	}))

	_, err := client.Holidays().Versions("hol1").Get(context.Background(), "missing")
	var nferr *NotFoundError
	if !errors.As(err, &nferr) {
		t.Fatalf("expected *NotFoundError, got %v (%T)", err, err)
	}
}

func TestHolidayVersions_Create(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/holidays/hol1/versions" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusCreated, "testdata/holiday_versions/get.json")
	}))

	v, err := client.Holidays().Versions("hol1").Create(context.Background(), CreateHolidayVersionParams{
		Name: "2024 Season",
		Code: "HV2024",
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if v.Code != "HV2024" {
		t.Fatalf("unexpected version: %+v", v)
	}
}

func TestHolidayVersions_Create_validationError(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusUnprocessableEntity, "testdata/holiday_versions/validation_error.json")
	}))

	_, err := client.Holidays().Versions("hol1").Create(context.Background(), CreateHolidayVersionParams{})
	var verr *ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("expected *ValidationError, got %v (%T)", err, err)
	}
}

func TestHolidayVersions_Update_nullableFields(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch || r.URL.Path != "/holidays/hol1/versions/ver1" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/holiday_versions/get.json")
	}))

	_, err := client.Holidays().Versions("hol1").Update(context.Background(), "ver1", UpdateHolidayVersionParams{
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

func TestHolidayVersions_Delete(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/holidays/hol1/versions/ver1" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	if err := client.Holidays().Versions("hol1").Delete(context.Background(), "ver1"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
}

func TestHolidayVersions_Restore(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/holidays/hol1/versions/ver1/restore" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/holiday_versions/get.json")
	}))

	v, err := client.Holidays().Versions("hol1").Restore(context.Background(), "ver1")
	if err != nil {
		t.Fatalf("Restore: %v", err)
	}
	if v.Code != "HV2024" {
		t.Fatalf("unexpected version: %+v", v)
	}
}
