---
name: update-tour-manager-go-client
description: >
  Update this hand-written Go client (RezKit Tour Manager, module
  github.com/rezkit/tour-manager-go) when openapi.yml changes, or when
  asked to add/fix a resource. Establishes the code-style conventions,
  the "clean Go over a messy spec" principle, and the breaking-change
  gate that keep the public API idiomatic and stable over time.
---

# Updating the Tour Manager Go client

This client is **hand-written, not generated**. `openapi.yml` in this repo is
the source of truth for what the API does, but the Go code is not derived
from it mechanically — every resource file imposes clean, idiomatic Go
conventions on top of whatever the spec documents, and sometimes has to
diverge from the spec's literal shape (see section 4). Read this whole
skill before making a change; it's short enough to read in full.

## 1. When to use this skill

- `openapi.yml` has been modified (new commit, PR, or manual edit) and the
  Go client needs to catch up.
- You're asked to "add resource X", "add endpoint Y", or "fix client gap
  for Z" — even without a spec diff in hand.
- You're asked to review whether the client is missing coverage for
  something the API now supports.

## 2. Diff the spec against current coverage

1. List every operation currently documented in the spec:
   ```
   grep -n 'operationId:' openapi.yml
   ```
   If you have a previous revision of the spec to compare against (e.g. a
   git ref before the change you're reacting to), diff the two:
   ```
   git diff <old-rev> -- openapi.yml
   ```
2. List every operation the client currently implements. Every exported
   method that maps to a spec operation carries a trailing
   `// spec: <operationId>` doc-comment line — this is the convention the
   whole client follows, so it's greppable:
   ```
   grep -rn '// spec:' *.go
   ```
3. Bucket the difference into: new operations (not yet implemented),
   changed operations (params, request/response schema, or status code
   changed), removed operations (client still implements something the
   spec no longer has), and unchanged.
4. For removed operations: this is very likely a breaking change (see
   section 5) — don't just delete the method quietly.

## 3. Code style rules

These are the conventions established by the framework and the three
reference resources (`holidays.go`, `categories.go`, `departures.go`).
Follow them exactly for consistency — don't invent a new shape per
resource.

- **Resource type**: `type <Plural>Resource struct { client *Client }`,
  obtained via `func (c *Client) <Plural>() *<Plural>Resource`. Fields are
  always unexported.
- **List + All pair**: every paginated list endpoint gets
  `List(ctx, opts <X>Options) (*Page[T], error)` for one page, and
  `All(ctx, opts <X>Options) iter.Seq2[T, error]` for auto-pagination,
  built via the shared `autoPaginate[T]` helper in `pagination.go`. See
  `holidays.go`'s `List`/`All` for the exact pattern to copy. Endpoints
  whose response is a plain `{"data": [...]}` with **no** pagination
  metadata get `List(ctx) ([]T, error)` only — no `Page[T]`, no `All`
  (an `All` that's always one call is misleading surface bloat). Use
  `fetchList[T]` from `pagination.go` for those.
- **Options structs field types** (apply consistently — this was a
  deliberate resolution of a real design ambiguity, not an accident):
  - `Page`/`Limit`: plain `int` via embedded `ListOptions` — 0 means
    unset, since valid pages/limits are always ≥ 1.
  - Free-text / sort / enum filters (`Search`, `Name`, `Sort`, `Order`,
    ...): plain `string` or a typed string enum — `""` means unset, since
    these types never have a legitimate empty value.
  - Filters backed by the API's `integer enum [0, 1]` convention
    (`Published`, `Trash`, `Searchable`, ...): `*bool` — `false` is a
    real, meaningful filter value distinct from "don't filter by this",
    so it needs three states (nil / false / true). Use `setBool01` from
    `query.go`.
  - A genuine boolean flag where "omitted" and "false" have the same
    effect (e.g. Categories' `children` — "include children" is either
    requested or it isn't): plain `bool`, no pointer. Use `setBool` from
    `query.go`.
- **Write params**: `Create<X>Params` (required fields as plain values,
  optional fields as pointers) and `Update<X>Params` (all fields
  pointers). Impose this convention even when the spec's request body is
  an untyped inline object (see section 4) — `categories.go` is the
  reference example of this. A field the spec marks `nullable: true` on a
  write body (distinguishing "don't touch" from "explicitly clear to
  null") uses `*Nullable[T]` (see `nullable.go`), not a plain `*T` — see
  `UpdateHolidayParams.Introduction`/`.Description` in `holidays.go`.
- **Attachment sub-resources**: reuse `AttachmentHandle[T]`
  (`attachment.go`) wrapped in a typed struct (see `CategoryAttachment` in
  `categories.go`) — never call `AttachmentHandle[T]` directly from
  outside the package's own resource files. Expose it two ways: a generic
  entry point on the attached resource
  (`client.Categories().For(entityType, id)`) and, where the parent
  resource has one, a same-shaped convenience method
  (`client.Holidays().Categories(id)`).
- **Restore**: use `restoreEmpty`/`restoreInto[T]` from `restore.go` —
  never add a third public generic `Restorable[T]` type. The public
  `Restore` method's signature on each resource must match exactly what
  the spec documents for that resource (some return the restored
  resource, most return nothing) — check both before assuming.
- **Discriminated unions** (`oneOf` schemas, with or without an explicit
  `discriminator:` in the spec): interface + concrete structs + a
  package-level `unmarshal<X>(raw) (X, error)` dispatcher keyed on the
  variant's own `type` field (or whatever field acts as an implicit
  discriminator if the spec has none — check every variant carries one
  reliably before assuming this pattern applies). See `inventory.go` for
  the full pattern, including how `MarshalJSON` injects the discriminator
  value on encode so callers never have to set it themselves.
- **Enums**: `EntityType` lives in `entitytype.go`; `SortOrder` and
  `OrderingCommand` live in `enums.go`. These are defined **once** —
  never redeclare them in a resource file. Per-resource sortable-field
  enums (`HolidaySortField`, `CategorySortField`, ...) follow the same
  `<Resource>Sort<Value>` prefix and live in the resource's own file. No
  exported mutable "allowed values" slices — use an unexported
  switch-based `Valid()` method if validation is useful.
- **Doc comments**: every exported identifier gets a real Go doc comment.
  This client's whole value proposition over the old generated client is
  developer experience — comments that read well in `go doc`/pkg.go.dev
  matter as much as the code.
- **`// spec: <operationId>` on every exported method** that maps to a
  spec operation — this is what section 2's coverage-diffing depends on.

## 4. The "impose clean conventions over a messy spec" principle

`openapi.yml` has real gaps and inconsistencies (untyped inline request
bodies, `required` fields missing from `properties`, endpoints with no
schema at all). Use this decision order:

- **OK to diverge from the spec's literal shape**: when the spec uses an
  untyped inline object with no named schema for a request body — impose
  the `Create<X>Params`/`Update<X>Params` convention regardless (see
  `categories.go`, where the spec's inline PATCH/POST bodies became clean
  typed params).
- **OK to include a `required`-but-undocumented field**: if a schema's
  `required` list references a field that `properties` never defines
  (confirmed spec bug, not absence), it's safe to conclude the field is
  real — include it, but comment exactly what's missing and why you're
  confident it exists (see `Holiday.Fields` in `holidays.go` for the
  template: it cites the exact spec inconsistency and ships as
  `json.RawMessage` rather than guessing a structured shape).
- **NOT OK to add a field or operation the spec doesn't confirm at all**
  (absent from both `required` and `properties`, or a path/method that
  doesn't exist in the spec) without independent corroborating evidence —
  and "the JS or PHP client has it" is not sufficient corroboration on its
  own, since those clients predate this spec and may themselves be ahead
  of or diverged from it. Track these as backlog items with a comment
  citing the gap; don't guess a shape and ship it silently. `holidays.go`
  and `categories.go` both have worked examples of this (`Copy` excluded
  entirely; `slug`/`seo`/`rank` excluded from `Holiday`; no
  single-category `Get`; no `ordering` on `UpdateCategoryParams`).
- **Never implement anything in the explicitly-out-of-scope set** (see the
  README's "Known gaps" section) until its spec gap is actually fixed —
  don't stub it with a guessed type "to get it working."
- When an endpoint's parameters are underdocumented but its response shape
  strongly implies a capability the spec forgot to declare (e.g.
  `listDepartures`'s response has full pagination metadata but its
  `parameters:` list omits `page`/`limit`), it's reasonable to infer the
  same convention used identically elsewhere in the spec applies — but
  say so in a comment (see `ListDeparturesOptions` in `departures.go`) so
  it's easy to find and confirm later.

## 5. Breaking-change handling

- This client's public API surface (every exported identifier) is subject
  to a breaking-change gate. Before and after your change, run:
  ```
  make apidiff
  ```
  which runs `gorelease` against the most recent git tag. Anything it
  reports as incompatible needs a deliberate decision, not an accident.
  **Known limitation**: `apidiff`/`gorelease`'s generics support has
  historically lagged the language. Changes touching `pagination.go`,
  `attachment.go`, or `nullable.go` (the generic-heavy files) should get a
  manual review of the public API even if `make apidiff` reports nothing,
  until this has been re-verified against a current toolchain version.
- Decision order when a spec change would naturally require a breaking Go
  change (a field becomes required, an enum value is removed, a response
  shape narrows):
  1. Can it be made additive instead (a new optional field, a new method,
     widening an enum)? Prefer this whenever it's semantically honest.
  2. If it's genuinely breaking because the **API itself** made a breaking
     change, the client should follow — ship it as its own clearly
     labeled commit/PR, not bundled with unrelated changes, and call out
     that it's a breaking release.
  3. Never break a signature just to "clean up" unrelated to an actual
     spec change — that's a separate, explicitly-approved change with its
     own justification.

## 6. Adding a new resource end-to-end

Use `holidays.go` (richest CRUD + a bare-array sub-resource + a
restore-returns-body variant), `categories.go` (attachment sub-resource +
messy-spec-to-clean-params), and `departures.go` (`*Params`/`*Properties`
triad + discriminated union) as templates, matching whichever pattern
your new resource actually needs:

1. Read the resource's paths and schemas in `openapi.yml` directly — grep
   for the schema name and every path referencing it. Identify: which of
   list/get/create/update/delete/restore exist, any non-CRUD actions,
   any sub-resources.
2. Write `List<X>Options` (embed `ListOptions`, add a `values() url.Values`
   method using the `query.go` helpers) and the model struct (apply
   section 4's principle for any spec gaps).
3. Write `Create<X>Params`/`Update<X>Params`.
4. Write `<X>Resource` with methods in this order:
   `List, All, Get, Create, Update, Delete, Restore`, each with a trailing
   `// spec: operationId` comment — omit any method whose operation
   doesn't exist in the spec (see `departures.go`'s missing `Restore`).
5. Wire any attachment sub-resources via `AttachmentHandle[T]` + a typed
   wrapper (section 3); wire any bespoke sub-resource (relations-shaped,
   not attachment-shaped) as its own small struct holding `client` plus
   the parent ID, in its own file (`holiday_relations.go` is the
   template).
6. Add the resource accessor to `Client`:
   `func (c *Client) <Plural>() *<Plural>Resource`.
7. Write httptest-based tests (section 8) and fixture JSON.
8. Run `go build ./...`, `go vet ./...`, `go test ./...`, `gofmt -l .`
   (must be empty) before considering the resource done.

## 7. Promoting a currently out-of-scope resource

When a resource listed in the README's "Known gaps" section gets a real
spec (both request AND response schemas — a partial fix, e.g. response
documented but request still ad hoc, still warrants the full "impose
clean Params" treatment from section 4, not a shortcut):

1. Follow the full "adding a new resource" checklist in section 6. Don't
   wire it in partially.
2. Remove it from the out-of-scope list in the README, and from any
   scope notes elsewhere (including this file if it's mentioned by name).

## 8. Testing requirements

- `httptest.Server`-based unit tests, one `_test.go` per resource file.
  Use the shared `newTestClient(t, handler) *Client` helper from
  `testhelpers_test.go`.
- Fixture JSON per case under `testdata/<resource>/<case>.json`, built to
  match the schema exactly — prefer copying a captured real response over
  hand-typing one, since hand-typed fixtures can accidentally validate a
  wrong assumption about the shape.
- Cover: the success path, at least one documented error status per
  operation (assert with `errors.As` against the specific typed error,
  e.g. `*tourmanager.NotFoundError`), and — for any field backed by a
  discriminated union or a `Nullable[T]` — a decode/encode test per
  variant/state (see `departures_test.go`'s inventory-variant tests and
  `holidays_test.go`'s `TestHolidaysResource_Update_nullableFields`).
- No live network calls, no real credentials, ever.
- Run `go build ./...`, `go vet ./...`, `gofmt -l .` (must print nothing),
  and `go test ./...` before considering any change complete.
