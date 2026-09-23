package tourmanager_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	tourmanager "github.com/rezkit/tour-manager-go"
)

// This example server stands in for the real Tour Manager API so the
// example is runnable/testable without network access.
func exampleServer() *httptest.Server {
	pages := map[string]string{
		"1": `{"total":2,"current_page":1,"last_page":2,"from":1,"to":1,"data":[{"id":"1","name":"Alpine Explorer","code":"ALP01"}]}`,
		"2": `{"total":2,"current_page":2,"last_page":2,"from":2,"to":2,"data":[{"id":"2","name":"Coastal Wander","code":"CST01"}]}`,
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		if page == "" {
			page = "1"
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, pages[page])
	}))
}

// Example_list shows fetching a single page of Holidays.
func Example_list() {
	srv := exampleServer()
	defer srv.Close()

	client := tourmanager.New("API_KEY", tourmanager.WithBaseURL(srv.URL))
	ctx := context.Background()

	page, err := client.Holidays().List(ctx, tourmanager.ListHolidaysOptions{
		ListOptions: tourmanager.ListOptions{Limit: 100, Page: 1},
	})
	if err != nil {
		panic(err)
	}
	for _, holiday := range page.Data {
		fmt.Printf("%s\t%s\n", holiday.Code, holiday.Name)
	}

	// Output:
	// ALP01	Alpine Explorer
}

// Example_all shows auto-paginating over every Holiday across pages.
func Example_all() {
	srv := exampleServer()
	defer srv.Close()

	client := tourmanager.New("API_KEY", tourmanager.WithBaseURL(srv.URL))
	ctx := context.Background()

	for holiday, err := range client.Holidays().All(ctx, tourmanager.ListHolidaysOptions{}) {
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\t%s\n", holiday.Code, holiday.Name)
	}

	// Output:
	// ALP01	Alpine Explorer
	// CST01	Coastal Wander
}
