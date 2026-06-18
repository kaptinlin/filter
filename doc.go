// Package filter provides headless value transformations for Go callers,
// template engines, and other dynamic-value runtimes.
//
// What it is: pure functions that take a runtime value and return a
// transformed value or an error. Strings, slices, maps, dates, and numbers.
//
// What it is not: a SQL builder, an ORM helper, an i18n framework, an XSS
// sanitizer, a query DSL, or a CLI argument parser. It does no I/O. It reads
// no environment variables. It does not consult a global locale or
// timezone. Functions that need "now" use SystemClock by default and expose
// Clock-injected variants for deterministic callers. Random and Shuffle are
// non-deterministic convenience filters; RandomWithRand and ShuffleWithRand
// accept explicit *rand.Rand values for deterministic callers.
//
// # Errors
//
// Every public function returns *Error or wraps one. Errors are categorized
// by Kind: KindInvalidInput, KindNotFound, KindArithmetic, KindFormat. Match
// with errors.Is(err, ErrNotFound) (matches by Kind).
//
// # Stability
//
// Stable semantics include UTF-8 length, UTC default timezone, nil/false-only
// falsiness, Find returning *Error{Kind:KindNotFound} on miss, Date's token
// grammar, and Map silently substituting nil for missing keys. See
// SPECS/00-runtime-contract.md for the durable runtime contract.
package filter
