# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with the filter package.

## Project Overview

**Module**: `github.com/kaptinlin/filter`
**Purpose**: Template filter library providing utilities for string manipulation, array operations, date formatting, number formatting, and mathematical computations.

This is a pure utility library with no state or configuration. All functions are designed to be used directly in template engines or as general-purpose data transformation utilities.

## Commands

```bash
# Testing
task test              # Run tests with race detection
go test -race ./...    # Direct test command

# Linting
task lint              # Run golangci-lint and go mod tidy checks
task golangci-lint     # Run only golangci-lint
task tidy-lint         # Verify go.mod/go.sum are clean

# Verification
task verify            # Full pipeline: deps, fmt, vet, lint, test
task fmt               # Format code
task vet               # Run go vet
task clean             # Remove ./bin directory

# Dependencies
task deps              # Download and tidy dependencies
task deps:update       # Update all dependencies
```

## Architecture

Single-package design organized by functional domain:

```
filter/
├── string.go          # String manipulation (trim, case, escape, encode, truncate, replace, slice, default)
├── array.go           # Slice operations (unique, unique-by, join, sort, where, find, compact, concat, random, shuffle)
├── time.go            # Date/time formatting via agentable/go-time, plus Clock-injected TimeAgoWithClock
├── number.go          # Number formatting (#,###.## DSL) and Bytes (SI units via agentable/go-humanize)
├── math.go            # Mathematical operations (abs, round, arithmetic)
├── data.go            # Nested data extraction with dot notation
├── rand.go            # SeededRand for deterministic RandomWithRand/ShuffleWithRand tests
├── truth.go           # Shared nil/false-only truthiness semantics
├── compare.go         # Shared numeric-first equality and ordering semantics
├── utils.go           # Type conversion utilities (toFloat64, toSlice)
├── errors.go          # *Error{Kind, Op, Path, Cause} model with four Kind sentinels
└── acronyms.go        # Acronym handling for case conversions
```

## SPECS Index

- [`SPECS/00-runtime-contract.md`](SPECS/00-runtime-contract.md) - Durable
  runtime contract, value semantics, consumer boundary, forbidden scope, and
  acceptance criteria.

## Key Types and Interfaces

This library uses primitive types and `any` for maximum flexibility:

```go
// All functions follow these patterns:

// String functions: string input/output
func Trim(input string) string
func Replace(input, old, replacement string) string
func Truncate(input string, maxLength int, ellipsis ...string) string

// HTML/encoding functions: string input/output
func Escape(input string) string
func EscapeOnce(input string) string
func Base64Encode(input string) string
func Base64Decode(input string) (string, error)

// Dual-mode functions: any input, any output
func Default(input, defaultValue any) any
func Slice(input any, offset int, length ...int) (any, error)

// Array functions: any input (slice/array), typed output
func Unique(input any) ([]any, error)
func UniqueBy(input any, key string) ([]any, error)
func SumBy(input any, key string) (float64, error)
func Random(input any) (any, error)
func RandomWithRand(r *rand.Rand, input any) (any, error)
func Shuffle(input any) ([]any, error)
func ShuffleWithRand(r *rand.Rand, input any) ([]any, error)
func Sort(input any, key ...string) ([]any, error)
func Where(input any, key string, value ...any) ([]any, error)
func Find(input any, key string, value any) (any, error)

// Data extraction: any input, any output
func Extract(input any, key string) (any, error)
```

### Error Types

All failures return `*Error{Kind, Op, Path, Cause}`. Callers branch on `Kind`
via `errors.Is` against four package-level sentinels (defined in `errors.go`):

- `ErrInvalidInput` (`KindInvalidInput`) - Wrong type, wrong shape, or out of range
- `ErrNotFound` (`KindNotFound`) - Missing path, key, or index in the input
- `ErrArithmetic` (`KindArithmetic`) - Division by zero, modulus by zero, etc.
- `ErrFormat` (`KindFormat`) - Parse failures (date layout, base64, URL escape, etc.)

`Op` and `Path` are diagnostic context only — `Is` matches by `Kind`. Build
errors through the internal helpers `invalidInput`, `notFound`, `arithmetic`,
and `formatErr`; never construct `*Error` literals at call sites.

## Coding Rules

### Zero-Panic Design
- **Never panic in production code** - all functions return errors
- Use custom error types from `errors.go` for consistent error handling
- Validate inputs and return descriptive errors

### UTF-8 Awareness
- Use `utf8.RuneCountInString()` for string length, not `len()`
- Use `[]rune` conversions for character-level operations
- Handle multi-byte characters correctly in truncation and case conversion

### Type Conversion Pattern
- Use `toSlice()` and `toFloat64()` utilities for consistent type handling
- Accept `any` for flexibility, validate types internally
- Return typed results when possible (e.g., `float64` for math operations)

### Go 1.26 Features
- Use `slices.Clone()`, `slices.Reverse()`, `slices.Max()`, `slices.Min()`
- Use `for i := range length` syntax
- Use `strings.Builder` with `Grow()` for efficient string building
- Use `math/rand/v2` for random operations

### Performance Patterns
- Pre-allocate slices with `make([]T, 0, capacity)` when size is known
- Use `strings.Builder.Grow()` for string concatenation
- Avoid unnecessary allocations in hot paths

## Testing

### Test Organization
- Each module has a corresponding `*_test.go` file
- `example_test.go` contains Godoc examples
- Use table-driven tests for comprehensive coverage

### Test Patterns
```go
// Use testify for assertions
func TestFunction(t *testing.T) {
    result, err := Function(input)
    require.NoError(t, err)
    assert.Equal(t, expected, result)
}

// Table-driven tests for multiple cases
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    any
        expected any
        wantErr  error
    }{
        // test cases
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test logic
        })
    }
}
```

### Running Tests
- Always run with race detection: `task test` or `go test -race ./...`
- Ensure 100% pass rate before committing
- Add tests for new functions and edge cases

## Dependencies

### Production Dependencies
- `github.com/agentable/go-time` - Date/time parsing and formatting (`time.go`)
- `github.com/agentable/go-humanize` - SI byte, number, and relative time formatting (`number.go`, `time.go`)
- `github.com/gosimple/slug` - URL slug generation for `Slugify()`
- `github.com/jinzhu/inflection` - Word pluralization for `Pluralize()`

### Test Dependencies
- `github.com/stretchr/testify` - Assertions and test utilities
- `github.com/google/go-cmp` - Deep comparison in table-driven tests

### Dependency Policy
- Minimize external dependencies
- Prefer stdlib when functionality overlaps
- Vet dependencies for maintenance and security

## Error Handling

### Pattern
```go
// Always return errors, never panic. Build errors through the internal
// helpers (invalidInput, notFound, arithmetic, formatErr) so that Op and
// Path are populated consistently.
func Function(input any) (result any, err error) {
    if invalid {
        return nil, invalidInput("Function", nil)
    }
    // process
    return result, nil
}

// Callers branch on Kind via the four sentinels.
if errors.Is(err, ErrNotFound) {
    // handle missing key/path/index
}
```

### Error Wrapping
- Wrap underlying causes through the helpers above; do not hand-roll `*Error`
- Map external library errors to a Kind (see `data.go:mapJSONPointerError`)
- Preserve cause chains via `Unwrap` so `errors.Is` keeps working

## Performance

### Benchmarking
- Add benchmarks for performance-critical functions
- Use `b.Loop()` (Go 1.24+) for benchmark loops
- Profile before optimizing

### Optimization Guidelines
- Pre-allocate slices and maps when size is predictable
- Use `strings.Builder` for string concatenation
- Avoid reflection in hot paths (use it only in `toSlice()` and `Size()`)
- Clone slices with `slices.Clone()` instead of manual copying


## Agent Skills

This package indexes agent skills from its own .agents/skills directory (filter/.agents/skills/):

| Skill | When to Use |
|-------|-------------|
| [agent-md-creating](.agents/skills/agent-md-creating/) | Create or update CLAUDE.md and AGENTS.md instructions for this Go package. |
| [code-simplifying](.agents/skills/code-simplifying/) | Refine recently changed Go code for clarity and consistency without behavior changes. |
| [committing](.agents/skills/committing/) | Prepare conventional commit messages for this Go package. |
| [dependency-selecting](.agents/skills/dependency-selecting/) | Evaluate and choose Go dependencies with alternatives and risk tradeoffs. |
| [go-best-practices](.agents/skills/go-best-practices/) | Apply Google Go style and architecture best practices to code changes. |
| [linting](.agents/skills/linting/) | Configure or run golangci-lint and fix lint issues in this package. |
| [modernizing](.agents/skills/modernizing/) | Adopt newer Go language and toolchain features safely. |
| [ralphy-initializing](.agents/skills/ralphy-initializing/) | Initialize or repair the .ralphy workflow configuration. |
| [ralphy-todo-creating](.agents/skills/ralphy-todo-creating/) | Generate or refine TODO tracking via the Ralphy workflow. |
| [readme-creating](.agents/skills/readme-creating/) | Create or rewrite README.md for this package. |
| [releasing](.agents/skills/releasing/) | Prepare release and semantic version workflows for this package. |
| [testing](.agents/skills/testing/) | Design or update tests (table-driven, fuzz, benchmark, and edge-case coverage). |
