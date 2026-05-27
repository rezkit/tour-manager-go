package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	tourmanager "github.com/rezkit/tour-manager"
)

func main() {
	ctx, done := context.WithTimeout(context.Background(), time.Minute)
	defer done()

	ctx = context.WithValue(ctx, tourmanager.ContextAccessToken, os.Getenv("TOUR_MANAGER_TOKEN"))

	client := tourmanager.NewAPIClient(tourmanager.NewConfiguration())

	result, _, err := client.DeparturesAPI.ListDepartures(ctx).Execute()
	if err != nil {
		fmt.Println("Failed to list departures", err)

		var e *tourmanager.GenericOpenAPIError = nil

		if errors.As(err, &e); e != nil {
			fmt.Println("Response body: ", string(e.Body()))
		}
	}

	fmt.Printf("Departures: %+v\n", result)
}
