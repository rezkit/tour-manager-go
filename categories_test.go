package tourmanager

import (
	"context"
	"net/http"
	"testing"
)

func TestCategoriesResource_List(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/holiday/categories" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/categories/list.json")
	}))

	page, err := client.Categories().List(context.Background(), EntityTypeHoliday, ListCategoriesOptions{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(page.Data) != 2 || page.Data[0].Name != "Skiing" {
		t.Fatalf("unexpected categories: %+v", page.Data)
	}
}

func TestCategoriesResource_List_queryParams(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if got := q.Get("children"); got != "true" {
			t.Errorf("children = %q, want %q", got, "true")
		}
		if got := q.Get("published"); got != "1" {
			t.Errorf("published = %q, want %q", got, "1")
		}
		if got := q.Get("searchable"); got != "0" {
			t.Errorf("searchable = %q, want %q", got, "0")
		}
		serveFile(t, w, http.StatusOK, "testdata/categories/list.json")
	}))

	published, searchable := true, false
	_, err := client.Categories().List(context.Background(), EntityTypeHoliday, ListCategoriesOptions{
		Children:   true,
		Published:  &published,
		Searchable: &searchable,
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
}

func TestCategoryAttachment_AttachReplaceDetach(t *testing.T) {
	var lastMethod string
	var lastQuery string
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantPath := "/holiday/hol123/categories"
		if r.URL.Path != wantPath {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		lastMethod = r.Method
		lastQuery = r.URL.RawQuery
		switch r.Method {
		case http.MethodPatch, http.MethodPut:
			serveFile(t, w, http.StatusOK, "testdata/categories/attach_response.json")
		case http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)
		}
	}))

	attachment := client.Categories().For(EntityTypeHoliday, "hol123")

	cats, err := attachment.Attach(context.Background(), []string{"cat1"})
	if err != nil {
		t.Fatalf("Attach: %v", err)
	}
	if lastMethod != http.MethodPatch || len(cats) != 1 {
		t.Fatalf("unexpected Attach result: method=%s cats=%+v", lastMethod, cats)
	}

	cats, err = attachment.Replace(context.Background(), []string{"cat1"})
	if err != nil {
		t.Fatalf("Replace: %v", err)
	}
	if lastMethod != http.MethodPut || len(cats) != 1 {
		t.Fatalf("unexpected Replace result: method=%s cats=%+v", lastMethod, cats)
	}

	if err := attachment.Detach(context.Background(), []string{"cat1", "cat2"}); err != nil {
		t.Fatalf("Detach: %v", err)
	}
	if lastMethod != http.MethodDelete {
		t.Fatalf("Detach used method %s", lastMethod)
	}
	if lastQuery != "ids=cat1&ids=cat2" {
		t.Fatalf("Detach query = %q", lastQuery)
	}
}

func TestHolidaysResource_Categories_delegatesToAttachmentHandle(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/holiday/hol123/categories" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/categories/list.json")
	}))

	page, err := client.Holidays().Categories("hol123").List(context.Background(), ListCategoriesOptions{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(page.Data) != 2 {
		t.Fatalf("unexpected categories: %+v", page.Data)
	}
}
