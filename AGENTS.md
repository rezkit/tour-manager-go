# AGENTS.md

Instructions for AI agents working with this repository — either as a
**consumer** importing this package into another project, or as a
**contributor** changing code in this repository itself. Read whichever
section applies; they're independent.

## What this is

`tour-manager-go` (module `github.com/rezkit/tour-manager-go`) is a
hand-written, idiomatic Go client for the RezKit Tour Manager API. It is
**not generated** from `openapi.yml` — every type and method is
deliberately hand-authored to be idiomatic Go, even where the spec's own
shape is looser or inconsistent. `openapi.yml` is the source of truth for
what the API *does*, not for what the Go code looks like.

Go 1.23+. Single flat package `tourmanager`. Zero third-party runtime
dependencies.

---

## Using this library

You're an agent writing code elsewhere that imports this package.

**Don't guess field or method names.** Every exported identifier has a
real doc comment written for exactly this purpose. Before writing code
against a type you haven't seen, run:

```sh
go doc github.com/rezkit/tour-manager-go
go doc github.com/rezkit/tour-manager-go.Holiday   # or whatever type
```

The patterns below are stable across the whole package; specific
field/method names differ per resource and must be looked up, not assumed
— and don't assume this client's shape matches the sibling `js`/`php`
clients' shape, they cover different scope (see "Known gaps" below).

### Construct a client

```go
client := tourmanager.New(apiKey)                                 // static API key, bearer auth
client := tourmanager.New(apiKey, tourmanager.WithBaseURL(url))    // e.g. a mock server in tests
client := tourmanager.New(apiKey, tourmanager.WithHTTPClient(hc))  // custom *http.Client (retries, tracing...)
```

`New` never returns an error and never panics — a bad config surfaces as a
normal `error` on the first call that needs it, not at construction time.

### Resources are `client.<Plural>()`

```go
client.Holidays()
client.Categories()
client.Departures()
client.Elements()
client.Prices()
client.Fields()
```

Each returns a small resource-scoped handle. ElementOptions has no
top-level accessor — it's reached via `client.Elements().Options(elementID)`
since the spec has no list operation for it. `Fields` manages custom-field
*definitions*, scoped by `EntityType` only (not a specific item — every
entity of a type shares the same definitions); the recorded *values* for
those fields live on the entity itself, typed as `CustomFieldsData` (see
`Holiday.Fields` for the one entity that currently exposes it).
`HolidayVersions` also has no top-level accessor — every operation on it
requires a holiday ID, so it's only reached via
`client.Holidays().Versions(holidayID)` (same shape as
`client.Holidays().Relations(id)`). Likewise `DepartureElements` — reached
via `client.Departures().Elements(departureID)` — has no top-level
accessor and no Create/Get/Delete/List: a Departure's elements are
generated automatically when an `Element` is attached and are only
discoverable via `Departure.Elements`; the only write operation is
`Update` (inventory and/or a balance-due override). **Only Holidays,
Categories, Departures, Elements, ElementOptions, Prices, Fields and
HolidayVersions are implemented as of this writing** — see "Known gaps"
below before assuming any other resource exists; check
`go doc github.com/rezkit/tour-manager-go` for the current accessor list on
`Client` if unsure.

### Pagination: `List` (one page) vs `All` (every page)

```go
page, err := client.Holidays().List(ctx, opts)   // *Page[Holiday]: Data, Total, CurrentPage, LastPage, ...

for holiday, err := range client.Holidays().All(ctx, opts) {  // iter.Seq2 — auto-paginates
    if err != nil {
        // an error can occur partway through iteration (e.g. page 3 of 5
        // failed) — always check it, don't just range and ignore
        break
    }
    _ = holiday
}
```

`opts` is a resource-specific `List<X>Options` struct embedding
`ListOptions{Page, Limit int}` — these are plain ints (0 = "use the API
default"), not pointers. A resource whose list endpoint has no pagination
metadata only exposes `List(ctx) ([]T, error)`, no `All`.

### Errors: `errors.As`, never string-match

```go
var verr *tourmanager.ValidationError
if errors.As(err, &verr) {
    // verr.Errors: map[string][]string, field name -> messages
}
```

Other typed errors, all satisfying `tourmanager.APIError` (`.StatusCode() int`):
`*NotFoundError` (404), `*GenericError` (400/409/429), `*UnauthorizedError`
(401), `*UnexpectedStatusError` (undocumented status). Transport-level,
non-API failures: `*RequestError` (no response reached), `*DecodeError`
(response didn't decode).

### Explicit null vs "leave unchanged" on updates

Some `Update<X>Params` fields are typed `*Nullable[T]`, not `*T` — the API
distinguishes "omit this field" (unchanged) from "set it to null"
(cleared) from "set it to a value", which a plain pointer can't express.

```go
params.Introduction = tourmanager.NullValue("new intro")  // set a value
params.Introduction = tourmanager.Null[string]()          // explicitly clear it
// params.Introduction left nil                            // leave unchanged
```

### Monetary values are `Decimal`, not `float64`

`Price.Value`/`Price.Deposit.Value` and `UpdatePriceParams`' equivalents
are `tourmanager.Decimal` (a defined string type), not `float64` — the API
encodes monetary amounts as JSON strings (e.g. `"1240.00"`) to avoid
floating-point rounding, despite `openapi.yml` documenting them as
`type: number`. `Decimal` has no arithmetic of its own; convert with a
decimal-math library or `strconv.ParseFloat` as your use case requires.
Don't assume every `type: number` field in the spec is actually a
`float64` in this client — check `go doc` for the field, not the spec.

### Attachment sub-resources (many-to-many, e.g. Categories on a Holiday)

```go
client.Holidays().Categories(holidayID)                            // sugar, where the parent resource has it
client.Categories().For(tourmanager.EntityTypeHoliday, holidayID)  // equivalent, generic form

attachment.Attach(ctx, ids)   // add, keep existing
attachment.Replace(ctx, ids)  // overwrite the whole set
attachment.Detach(ctx, ids)   // remove
attachment.List(ctx, opts)    // as above
```

### Runnable reference code

[`examples/`](examples/README.md) has small, complete `package main`
programs demonstrating these patterns end to end (e.g. `list-holidays` for
`All` + `text/tabwriter`). Prefer copying from there over inferring usage
from this file alone — the examples are compiled and tested, this file
isn't.

### Don't

- Don't assume every field is a pointer — this isn't the old
  openapi-generator client it replaced. Required fields are plain values;
  only genuinely optional or tri-state fields are pointers.
- Don't invent a resource or method "by analogy" with the JS/PHP sibling
  clients. This client's scope is deliberately smaller right now (see
  "Known gaps"). Missing means not implemented yet, not a bug in your call.
- Don't omit `ctx`/pass `context.Background()` reflexively in code meant
  for someone else to use — every call takes it as the bounding mechanism;
  there's no built-in request timeout.

### Known gaps

Only **Holidays**, **Categories**, **Departures**, **Elements**,
**ElementOptions**, **Prices**, **Fields**, **HolidayVersions** are
implemented. Beyond that, some operations are deliberately excluded even on
implemented
resources because `openapi.yml` doesn't document them confidently enough
to ship (e.g. Holiday has no `Copy`, Categories has no single-item `Get`).
Full list and reasoning: see `README.md`'s "Known gaps / out of scope"
section.

---

## Contributing to this library

You're an agent changing code in this repository.

**Read `.claude/skills/update-tour-manager-go-client/SKILL.md` before
making any change.** It's the authoritative, detailed process for this
repo — spec-diffing against `openapi.yml`, the full code-style rulebook,
the "clean Go over a messy spec" decision process, breaking-change
handling, and a checklist for adding a new resource. It's plain Markdown;
read and follow it directly even if your tooling has no notion of
"skills" — nothing about it is Claude-Code-specific except its file
location.

The essentials, if you read nothing else:

1. `openapi.yml` is the source of truth for API *behavior*, but the Go
   code is never derived from it mechanically — impose the repo's
   established conventions even where the spec itself is loose or
   inconsistent (skill, section 3–4).
2. Every exported method that maps to a spec operation carries a trailing
   `// spec: <operationId>` doc comment. This is how coverage is tracked:
   `grep -rn '// spec:' *.go`.
3. Use `holidays.go`, `categories.go`, `departures.go` as templates —
   respectively: full CRUD plus a bare-array sub-resource, an attachment
   sub-resource, and the create/update Params triad plus a discriminated
   union. The shared plumbing they build on lives in `client.go`,
   `transport.go`, `errors.go`, `pagination.go`, `attachment.go`,
   `nullable.go`, `entitytype.go`, `enums.go`, `restore.go`, `query.go` —
   read those before adding a new generic abstraction, the piece you need
   may already exist.
4. Before considering any change done:
   ```sh
   go build ./... && go vet ./... && gofmt -l . && go test ./...
   ```
   (`gofmt -l .` must print nothing). Run `make apidiff` before shipping
   anything that touches an exported identifier, to check the public API
   surface against the last tag.
5. Never guess an undocumented request/response shape and ship it as if
   confirmed. Track it as an explicit gap instead (see `README.md`'s
   "Known gaps" section and the skill's section 4 for the exact decision
   rule) — "the JS/PHP client has it" is not sufficient evidence on its
   own, those clients may be ahead of or diverged from this spec.
6. Never make a public API change (removed/renamed exported identifier,
   changed method signature, narrowed field type) without treating it as
   a deliberate breaking change, flagged as such — see the skill's
   section 5.
