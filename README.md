# tour-manager-go

A hand-written, idiomatic Go client for the [RezKit Tour Manager](https://rezkit.app)
API.

This client is built by hand against `openapi.yml` (the source of truth for
API behavior), not generated from it — see
`.claude/skills/update-tour-manager-go-client/SKILL.md` for the conventions
and process used to keep it that way as the API evolves.

## On the use of language models

This client's code was designed and written with substantial assistance
from LLM-based coding agents, and that process is expected to continue as
the API evolves — that's what `AGENTS.md` and the update skill referenced
above exist to steer consistently.

The reason isn't speed for its own sake. A purely scripted generator (such
as the `openapi-generator`-based client this one replaced) can only ever
transform the spec mechanically: every schema becomes a struct, every
optional field becomes a pointer, every unnamed inline body becomes a
throwaway type like `InlineObject3`. It has no way to notice that a dozen
endpoints share an attachment pattern that deserves one shared abstraction,
or that a spec's own request shape is worth diverging from in favor of a
cleaner Go convention. An LLM-based process can read the spec the way a
human maintainer would — semantically, not mechanically — and make exactly
those judgment calls, while still being systematic enough that the result
is applied consistently across the whole client rather than ad hoc. The
skill file is what keeps it systematic: the conventions, and the reasoning
behind every deliberate deviation from the spec's literal shape, are
written down and re-applied on every update rather than reinvented.

None of this changes normal engineering standards: every change is
reviewed, and is expected to pass `go build`, `go vet`, `gofmt`, `go test`,
and an API-surface diff before it ships (see "Development" below).

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

### Monetary values

`Price.Value`/`Price.Deposit.Value` (and `UpdatePriceParams`' equivalents)
are `tourmanager.Decimal`, not `float64`. `openapi.yml` documents them as
`type: number`, but the API actually encodes them as JSON strings (e.g.
`"1240.00"`) to avoid floating-point rounding on currency values — a
spec/wire mismatch confirmed against real API behavior, not a client guess.
`Decimal` is a defined string type with no arithmetic of its own; convert
with a decimal-math library of your choice, or `strconv.ParseFloat` if
approximate float precision is acceptable for your use case.

## Examples

Runnable sample programs live under [`examples/`](examples/README.md), one
per directory — e.g. `REZKIT_API_KEY=... go run ./examples/list-holidays`.

## Known gaps / out of scope

This client currently implements eight fully-worked resources — **Holidays**,
**Categories**, **Departures**, **Elements**, **ElementOptions**,
**Prices**, **Fields** and **HolidayVersions** — chosen to exercise every
framework pattern (pagination, attachment, the create/update params
convention, discriminated unions, and a bespoke CRUD sub-resource shaped
like `HolidayRelations`, which `HolidayVersions` itself now also follows).
`Departures` also exposes `DepartureElements` (`client.Departures().Elements(id)`),
a single-method (`Update` only) sub-resource for setting a
`DepartureElement`'s inventory and balance-due override, closing the
"discovering generated IDs" and "no write endpoint" gaps from the original
coverage report together. `CustomFieldsData` (the recorded values for a
`Fields`-defined field,
distinct from the definitions themselves) is standalone, reusable infra —
currently wired into `Holiday.Fields`, and ready for `Location`,
`Accommodation` and `Extra` to pick up once those resources are built,
since all three reference the same schema. The remaining ~20 resources in
`openapi.yml` are an explicit backlog: follow
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
  operation for this resource. Conversely, `CreateDepartureParams.VersionID`
  is included despite being completely undocumented in openapi.yml's
  `DepartureProperties` schema (absent from both `properties` and
  `required`) — without it, `createDeparture` cannot actually associate the
  new Departure with a Holiday Version at all, confirmed against real API
  behavior rather than inferred from the spec (see AGENTS.md's note on
  `openapi.yml` being agent-generated from the implementation, not
  hand-authored).
- **Elements**: no `Restore` method, for the same reason.
- **Fields**: `Update` only supports Text and Number fields
  (`UpdateTextFieldParams`/`UpdateNumberFieldParams`) — openapi.yml's
  `updateField` request body documents just those two of the five field
  types; Date, Boolean and Selection fields have no update body schema at
  all. No `Get`/`Delete`/`Restore` for fields or groups either — none of
  those operations exist in the spec.
- **HolidayVersion**: no `deleted_at` field — openapi.yml's `HolidayVersion`
  schema has no `deleted_at` property (and no `required` list at all),
  despite `Delete`/`Restore` both existing. `UpdateHolidayVersionParams`
  also excludes `fields`/`ordering`, even though `updateHolidayVersion`
  reuses `UpdateHoliday`'s request body verbatim (which documents both) —
  neither has any corresponding property on `HolidayVersion`'s own read
  schema, so they're treated as artifacts of the reused body rather than
  confirmed capabilities, the same reasoning as Holiday's excluded
  `slug`/`seo`/`rank`.

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

## License

MIT — see [LICENSE](LICENSE).
