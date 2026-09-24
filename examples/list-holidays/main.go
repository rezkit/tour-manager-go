// Command list-holidays prints every Holiday in the account as an aligned
// table, using client.Holidays().All to auto-paginate and text/tabwriter
// for the tabular output.
//
// Usage:
//
//	REZKIT_API_KEY=... go run ./examples/list-holidays
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	tourmanager "github.com/rezkit/tour-manager-go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	apiKey := os.Getenv("REZKIT_API_KEY")
	if apiKey == "" {
		return errors.New("REZKIT_API_KEY environment variable must be set")
	}

	client := tourmanager.New(apiKey)

	// Bound every call with a context, since the client has no built-in
	// request timeout of its own.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "CODE\tNAME\tPUBLISHED\tCREATED")

	for holiday, err := range client.Holidays().All(ctx, tourmanager.ListHolidaysOptions{}) {
		if err != nil {
			// Flush whatever rows already printed before reporting the
			// error that ended iteration early.
			_ = tw.Flush()
			return fmt.Errorf("listing holidays: %w", err)
		}
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n",
			holiday.Code,
			holiday.Name,
			publishedLabel(holiday.Published),
			holiday.CreatedAt.Format("2006-01-02"),
		)
	}

	return tw.Flush()
}

func publishedLabel(published bool) string {
	if published {
		return "yes"
	}
	return "no"
}
