// Package tourmanager is a hand-written, idiomatic Go client for the
// RezKit Tour Manager API.
//
// Create a [Client] with [New] and call one of its resource accessors,
// such as [Client.Holidays]:
//
//	client := tourmanager.New(os.Getenv("REZKIT_API_KEY"))
//
//	page, err := client.Holidays().List(ctx, tourmanager.ListHolidaysOptions{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, h := range page.Data {
//		fmt.Println(h.Code, h.Name)
//	}
//
// Every list endpoint that supports pagination also has an All method that
// returns an iterator over every matching item, fetching subsequent pages
// as needed:
//
//	for holiday, err := range client.Holidays().All(ctx, tourmanager.ListHolidaysOptions{}) {
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(holiday.Code, holiday.Name)
//	}
//
// Errors returned by API calls implement [APIError]; use [errors.As] to
// inspect a specific error, such as [ValidationError]:
//
//	var verr *tourmanager.ValidationError
//	if errors.As(err, &verr) {
//		for field, messages := range verr.Errors {
//			fmt.Println(field, messages)
//		}
//	}
//
// # Coverage
//
// This client does not yet cover the full Tour Manager API surface. See
// the "Known gaps" section of the repository README, and the
// update-tour-manager-go-client skill, for what's implemented and what's
// intentionally deferred until the OpenAPI spec documents it.
package tourmanager
