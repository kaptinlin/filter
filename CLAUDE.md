# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with the filter package.

## Project Overview

**Module**: `github.com/kaptinlin/filter`
**Go Version**: 1.26
**Purpose**: Template filter library providing utilities for string manipulation, array operations, date formatting, number formatting, and mathematical computations.

This is a pure utility library with no state or configuration. All functions are designed to be used directly in template engines or as general-purpose data transformation utilities.

## Commands

```bash
# Testing
task test              # Run tests with race detection
go test -race ./...    # Direct test command

# Linting
task lint              # Run golangci-lint and go mod tidy checks
make golangci-lint     # Run only golangci-lint
make tidy-lint         # Verify go.mod/go.sum are clean

# Build
make all               # Run lint + test
task clean             # Remove ./bin directory

# Dependencies
task deps              # Download and tidy dependencies
```

## Architecture

Single-package design organized by functional domain:

```
filter/
├── string.go          # String manipulation (trim, case conversion, pluralization, slugify)
├── array.go           # Slice operations (unique, join, shuffle, aggregations)
├── date.go            # Date/time formatting using carbon library
├── number.go          # Number formatting (including byte formatting)
├── math.go            # Mathematical operations (abs, round, arithmetic)
├── data.go            # Nested data extraction with dot notation
├── utils.go           # Type conversion utilities (toFloat64, toSlice)
├── errors.go          # Centralized error definitions
└── acronyms.go        # Acronym handling for case conversions
```

## Key Types and Interfaces

This library uses primitive types and `any` for maximum flexibility:

```go
// All functions follow these patterns:

// String functions: string input/output
func Trim(input string) string
func Replace(input, old, replacement string) string

// Array functions: any input (slice/array), typed output
func Unique(input any) ([]any, error)
func Max(input any) (float64, error)

// Data extraction: any input, any output
func Extract(input any, key string) (any, error)
```

### Error Types

All errors are defined in `errors.go`:

- `ErrNotNumeric` - Non-numeric input to numeric functions
- `ErrInvalidTimeFormat` - Invalid time parsing
- `ErrUnsupportedType` - Unsupported data types
- `ErrNotSlice` - Non-slice input to slice functions
- `ErrEmptySlice` - Operations on empty slices
- `ErrInvalidArguments` - Invalid function arguments
- `ErrKeyNotFound` - Missing keys in data extraction
- `ErrIndexOutOfRange` - Invalid array/slice indices
- `ErrInvalidKeyType` - Invalid key types in data extraction
- `ErrDivisionByZero` - Division by zero
- `ErrModulusByZero` - Modulus by zero
- `ErrUnsupportedSizeType` - Unsupported types in Size filter
- `ErrNegativeValue` - Non-negative value requirements

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
- `github.com/dromara/carbon/v2` - Date/time handling with rich formatting
- `github.com/kaptinlin/jsonpointer` - JSON pointer traversal for `Extract()`
- `github.com/gosimple/slug` - URL slug generation for `Slugify()`
- `github.com/jinzhu/inflection` - Word pluralization for `Pluralize()`
- `github.com/dustin/go-humanize` - Human-readable formatting for `Bytes()`

### Test Dependencies
- `github.com/stretchr/testify` - Assertions and test utilities

### Dependency Policy
- Minimize external dependencies
- Prefer stdlib when functionality overlaps
- Vet dependencies for maintenance and security

## Error Handling

### Pattern
```go
// Always return errors, never panic
func Function(input any) (result any, err error) {
    if invalid {
        return nil, ErrInvalidArguments
    }
    // process
    return result, nil
}

// Use errors.Is() for error checking
if errors.Is(err, ErrKeyNotFound) {
    // handle missing key
}
```

### Error Wrapping
- Wrap errors with context using `fmt.Errorf("%w: context", err)`
- Map external library errors to filter errors (see `data.go:mapJSONPointerError`)
- Preserve error chains for debugging

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
