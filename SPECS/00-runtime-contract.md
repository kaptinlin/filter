# Runtime Contract

## Overview

`filter` is a headless value-transformation library. Public functions accept
ordinary Go values and return transformed values or typed errors. The package is
safe to embed in higher-level runtimes because it owns runtime value semantics
without owning parsing, rendering, registries, trusted-string policy, source
locations, localization, or application configuration.

> **Why**: The most common consumer is an adapter that already has its own
> syntax, error presentation, truthiness for control flow, and rendering-safety
> rules. `filter` stays useful for that consumer by being small, pure, and
> predictable.
>
> **Rejected**: Matching another product's behavior by name was rejected because
> it makes correctness depend on an external moving target. A broad internal
> value object was rejected because the implemented call sites are clearer with
> small named functions. Global switches were rejected because they make the same
> function mean different things in different processes.
>
> **Basis**: Current implementation and tests in this repository, plus the
> downstream adapter test run against `/Users/kaptinlin/code/golang/template`.

## Concept Model

### Headless Transformation

- **Definition**: A pure function from value plus arguments to value or error.
- **Includes**: string, collection, data lookup, date, number, math, and random
  convenience filters.
- **Excludes**: rendering, escaping policy beyond explicit string helpers,
  parser state, environment lookup, I/O, registries, locale policy, and
  timezone configuration.
- **Invariant**: No public function reads global application state except
  deterministic clock and random defaults documented by this package.

### Dynamic Value

- **Definition**: Any Go value accepted at a public boundary.
- **Owner**: This spec owns package-level value semantics; callers own any
  higher-level language semantics they layer above these functions.
- **Invariant**: Shared semantics are named once in implementation and tested at
  public boundaries: path lookup, missing policy, truthiness, numeric coercion,
  comparison, sizing, slicing, and formatting.

### Path

- **Definition**: A compact accessor grammar for extracting nested values.
- **Grammar**: `.` separates segments, `\.` is a literal dot, `\\` is a literal
  backslash, and decimal segments index slices or arrays.
- **Includes**: maps, slices, arrays, structs, pointers, and interfaces.
- **Excludes**: query languages, wildcards, filters, predicates, and recursive
  traversal.
- **Invariant**: Path-based failures populate `*Error.Path` with the original
  accessor string.

## Runtime Contracts

### Error Contract

- Public failures return or wrap `*Error`.
- `errors.Is` matches by one of four sentinels: `ErrInvalidInput`,
  `ErrNotFound`, `ErrArithmetic`, or `ErrFormat`.
- `Op` and `Path` are diagnostic context; callers must branch on the sentinel,
  not on message text.
- External parsing or conversion failures are mapped into the four package
  kinds while preserving cause chains.

### Lookup And Missing Policy

`Extract` is the public translation layer over internal path lookup. Collection
filters choose explicit missing policies instead of relying on accidental
lookup errors.

| Filter | Missing or unreachable path |
|---|---|
| `Map` | preserves input length and substitutes `nil` |
| `Sort`, `SortNatural` | sorts the item as though the key were `nil` |
| `Compact` | skips the item when the requested key is missing or `nil` |
| `Where`, `Reject`, `Has` | treats missing as no match |
| `Find` | skips missing values and returns `ErrNotFound` when no item matches |
| `FindIndex` | skips missing values and returns `-1` when no item matches |
| `UniqueBy` | returns an error |
| `SumBy` | returns an error |

> **Why**: Different filters naturally answer different caller questions. A map
> operation should preserve shape, a predicate should filter, and an aggregate
> should fail when its input is incomplete.
>
> **Rejected**: A global strictness mode was rejected because it would make
> local behavior depend on hidden process state. A single "missing is nil"
> policy was rejected because it would hide incomplete aggregate input.
>
> **Basis**: `array_test.go` covers the policy table through public filters.

### Truthiness And Defaulting

- Only `nil` and `false` are falsy.
- Empty strings, zero numbers, empty slices, and empty maps are truthy.
- `Default` uses the same falsy rule as predicate filters.
- Predicate filters without an explicit comparison value use this truthiness
  rule.

### Equality And Ordering

- Equality compares numeric Go values and decimal numeric strings by numeric
  value before falling back to ordinary or deep equality.
- Ordering compares numeric Go values and decimal numeric strings numerically
  before falling back to string comparison.
- Natural ordering uses the same numeric-first rule and applies
  case-insensitive string fallback.
- `nil` sorts before non-`nil`.

> **Why**: Higher-level callers often pass loosely typed data. Numeric-first
> comparison makes values that are semantically numbers behave as numbers while
> keeping non-numeric values simple.
>
> **Rejected**: String-only comparison was rejected because it sorts numeric
> values incorrectly. A coercion framework was rejected because current callers
> need only equality, ordering, and numeric conversion primitives.
>
> **Basis**: `TestCollectionNumericEquality`, `TestSort`, `TestSortNatural`,
> and `TestUniqueByCrossTypeNumericKey`.

### Size, Slice, And Numeric Conversion

- String length and string slicing are UTF-8 rune based.
- Slice and array slicing use element offsets.
- Negative offsets count from the end.
- Out-of-range, zero-length, and negative-length slices return an empty value
  rather than panicking.
- Numeric aggregate filters coerce numeric Go values and decimal strings.
- Exact-integer conversion accepts only values that can be represented as an
  `int64` without fractional loss or overflow.
- `Bytes` accepts only non-negative whole-number byte counts.

### Formatting

- `Date` defaults parsed date strings to UTC.
- `Date` owns a stable token grammar; unknown letters pass through as literals,
  and a backslash escapes the next byte.
- `Number` owns a compact `#,###.##`-style grammar: decimal precision is
  derived from characters after `.`, and `,` in the integer part enables
  grouping.
- Non-finite number formatting is explicit: `NaN`, `+Inf`, and `-Inf` render as
  tokens.
- Formatting functions do not own locale, translation, or timezone policy.

### Randomness

- `Random` and `Shuffle` are non-cryptographic convenience helpers.
- Deterministic callers use `SeededRand` with `RandomWithRand` or
  `ShuffleWithRand`.
- Passing a nil random source to deterministic variants returns
  `ErrInvalidInput`.

## Consumer Boundary

Consumers own:

- filter names and aliases
- argument defaults and arity presentation
- source-position errors
- rendering and trusted-string policy
- escaping strategy
- control-flow truthiness outside this package
- localization and timezone policy
- application configuration and dependency injection

This package owns only the behavior of its exported Go functions.

> **Why**: The boundary keeps the library stable for more than one runtime. It
> also prevents higher-level concerns from forcing additional global state into
> a pure utility package.
>
> **Rejected**: Moving consumer registries, trusted string types, or source
> locations into `filter` was rejected because those concepts have no meaning
> without the consuming runtime.
>
> **Basis**: Public API shape, `doc.go`, and the downstream adapter test run.

## Forbidden

- Do not define behavior by naming a reference product. State the rule directly
  and test it locally.
- Do not add global strictness, locale, timezone, registry, or safety switches.
- Do not add compatibility shims whose only purpose is preserving weaker old
  behavior.
- Do not add query-language features to path lookup.
- Do not introduce an all-purpose internal value wrapper unless it removes more
  code than it adds across real call sites.
- Do not put consumer syntax, rendering, source locations, or trusted-string
  policy into this package.
- Do not duplicate this contract in README-style prose; behavior belongs in
  tests, and this spec is the durable map.

## Acceptance Criteria

- `task verify` passes for this repository.
- The downstream adapter repository passes `go test ./...` while consuming this
  local checkout through a temporary Go workspace.
- Public behavior tests cover path lookup, path error diagnostics, collection
  missing policies, truthiness, defaulting, numeric equality, numeric ordering,
  exact integer conversion, string slicing, formatter grammars, and random
  determinism.
- `SPECS/` does not name reference products as behavioral authorities.
- There is no root-level improvement plan carrying completed implementation
  guidance after this spec is updated.

**Origin:** migrated from the former root runtime contract.
