package tourmanager

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestPricesResource_Get(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/holidays/prices/01gpkgcy6t0m84czh8gy4kprc1" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		serveFile(t, w, http.StatusOK, "testdata/prices/get.json")
	}))

	p, err := client.Prices().Get(context.Background(), "01gpkgcy6t0m84czh8gy4kprc1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if p.Currency != "USD" || p.Value != 1240.0 {
		t.Fatalf("unexpected price: %+v", p)
	}
	if p.Deposit == nil || p.Deposit.Value != 200.0 {
		t.Fatalf("unexpected deposit: %+v", p.Deposit)
	}
	if !p.Initialized {
		t.Errorf("Initialized = false, want true")
	}
}

func TestPricesResource_Get_notFound(t *testing.T) {
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFile(t, w, http.StatusNotFound, "testdata/prices/not_found.json")
	}))

	_, err := client.Prices().Get(context.Background(), "missing")
	var nferr *NotFoundError
	if !errors.As(err, &nferr) {
		t.Fatalf("expected *NotFoundError, got %v (%T)", err, err)
	}
}

func TestPricesResource_Update_clearsDeposit(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch || r.URL.Path != "/holidays/prices/prc1" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/prices/get.json")
	}))

	_, err := client.Prices().Update(context.Background(), "prc1", UpdatePriceParams{
		Deposit: ptr(Null[float64]()),
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	body := string(gotBody)
	if !strings.Contains(body, `"deposit":null`) {
		t.Errorf("expected explicit null deposit: %s", body)
	}
	if strings.Contains(body, `"value"`) || strings.Contains(body, `"on_sale"`) {
		t.Errorf("expected value/on_sale to be omitted (unset): %s", body)
	}
}

func TestPricesResource_Update_setsValue(t *testing.T) {
	var gotBody []byte
	client := newTestClient(t, handlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		gotBody, err = io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		serveFile(t, w, http.StatusAccepted, "testdata/prices/get.json")
	}))

	value := 999.5
	_, err := client.Prices().Update(context.Background(), "prc1", UpdatePriceParams{
		Value: &value,
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	if body := string(gotBody); !strings.Contains(body, `"value":999.5`) {
		t.Errorf("unexpected request body: %s", body)
	}
}
