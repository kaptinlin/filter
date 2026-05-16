# Public Contract

`filter` is a headless value-transformation package. It does no I/O, reads no
environment variables, and has no global locale, timezone, registry, or
strictness setting.

## Stable Semantics

- Errors returned by this package match one of four sentinels:
  `ErrInvalidInput`, `ErrNotFound`, `ErrArithmetic`, or `ErrFormat`.
- String length and string slicing are UTF-8 rune based, not byte based.
- Date string parsing defaults to UTC. `Date` uses PHP/Carbon-style tokens,
  not Go reference-time layouts.
- Relative time uses `SystemClock` through `TimeAgo`, or an injected `Clock`
  through `TimeAgoWithClock`.
- Predicate truthiness treats only `nil` and `false` as falsy. Empty strings,
  zero numbers, and empty collections are truthy.
- `Default` falls back only for `nil` and `false`; `""`, `0`, and empty
  collections are caller-supplied values and are returned unchanged.
- `Extract` uses a small dot-path grammar: `.` separates segments, `\.` is a
  literal dot, `\\` is a literal backslash, and decimal segments index
  slices/arrays. It does not implement JSONPath.
- `Map` preserves input length and substitutes `nil` when a key cannot be
  extracted from an element.
- `Find` returns `ErrNotFound` when no item matches.
- `Unique` requires comparable elements. `UniqueBy` supports property-based
  uniqueness, including non-comparable extracted values.
- `Bytes` accepts only non-negative whole-number byte counts.
- `Random` and `Shuffle` are non-cryptographic convenience helpers.
  Deterministic callers pass `SeededRand` to `RandomWithRand` or
  `ShuffleWithRand`.

## Consumer Boundary

Consumers own aliases, argument defaults, template syntax, rendering safety,
trusted string types, localization policy, error presentation, and application
configuration. Behavior that depends on those concerns stays outside this
package.
