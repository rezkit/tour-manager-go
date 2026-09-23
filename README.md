# tour-manager-go

A hand-written, idiomatic Go client for the [RezKit Tour Manager](https://rezkit.app)
API.

This client is built by hand against `openapi.yml` (the source of truth for
API behavior), not generated from it — see
`.claude/skills/update-tour-manager-go-client/SKILL.md` for the conventions
and process used to keep it that way as the API evolves.

## Installation

```sh
go get github.com/rezkit/tour-manager-go
```

Requires Go 1.23 or later. This module has zero third-party runtime
dependencies.

## Usage

```go
import tourmanager "github.com/rezkit/tour-manager-go"

client := tourmanager.New(os.Getenv("REZKIT_API_KEY"))
ctx := context.Background()

// Fetch a single page.
page, err := client.Holidays().List(ctx, tourmanager.ListHolidaysOptions{
	ListOptions: tourmanager.ListOptions{Limit: 100, Page: 1},
})
if err != nil {
	log.Fatal(err)
}
for _, holiday := range page.Data {
	fmt.Println(holiday.Code, holiday.Name)
}

// Or auto-paginate across every page.
for holiday, err := range client.Holidays().All(ctx, tourmanager.ListHolidaysOptions{}) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(holiday.Code, holiday.Name)
}
```

### Errors

Every non-2xx response is returned as an error implementing
`tourmanager.APIError`. Use `errors.As` to inspect a specific kind, such as
field-validation failures:

```go
_, err := client.Holidays().Create(ctx, tourmanager.CreateHolidayParams{Name: "x", Code: "y"})

var verr *tourmanager.ValidationError
if errors.As(err, &verr) {
	for field, messages := range verr.Errors {
		fmt.Println(field, messages)
	}
}
```

### Attaching sub-resources

Resources like Categories can attach to multiple kinds of parent entity.
Access an item's attached Categories either generically or via a parent
resource's convenience method:

```go
client.Categories().For(tourmanager.EntityTypeHoliday, holidayID)
client.Holidays().Categories(holidayID) // equivalent
```

```go
attachment := client.Holidays().Categories(holidayID)
_, err := attachment.Attach(ctx, []string{categoryID})   // adds, keeps existing
_, err  = attachment.Replace(ctx, []string{categoryID})  // overwrites
err     = attachment.Detach(ctx, []string{categoryID})   // removes
```

## Known gaps / out of scope

This client currently implements three fully-worked resources — **Holidays**,
**Categories** and **Departures** — chosen to exercise every framework
pattern (pagination, attachment, the create/update params convention,
discriminated unions). The remaining ~25 resources in `openapi.yml` are an
explicit backlog: follow
`.claude/skills/update-tour-manager-go-client/SKILL.md` to add them.

Separately, the following are excluded from every resource until
`openapi.yml` documents them properly — they are not guessed at:

- **No schema anywhere in the spec**: MapLines, MapMarkers, RoomTypes,
  RoomTypePrices.
- **Status code documented, but no request/response body schema**: Content
  item CRUD, Images/Locations attachment, Cake/Extra/Accommodation update
  bodies, `searchHolidays`/`suggestHolidays`/`reindexHolidays`.
- **Holiday**: `Copy` (response documented, request body isn't), and the
  `slug`/`seo`/`rank` fields (referenced by the JS client but absent from
  the `Holiday` schema in `openapi.yml`).
- **Categories**: a single-category `Get` (no such path/operation exists in
  the spec), and an `ordering`/reorder field on `UpdateCategoryParams` (the
  spec's inline update body doesn't include one, unlike Holiday's).
- **Departures**: no `Restore` method — the spec defines no restore
  operation for this resource.

See the skill file for the full decision process behind these exclusions,
and how to promote a resource out of this list once its spec gap is fixed.

## Development

```sh
make build   # go build ./...
make test    # go test ./...
make vet     # go vet ./...
make fmt     # gofmt -l . (should print nothing)
make apidiff # check exported API surface against the last git tag
```
