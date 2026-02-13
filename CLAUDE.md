# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Testing
```bash
make test        # Run all tests with race detection
go test -race ./... # Alternative test command
```

### Linting and Code Quality
```bash
make lint        # Run linting (includes golangci-lint and go mod tidy)
make golangci-lint # Run only golangci-lint
make tidy-lint   # Check go.mod/go.sum are up to date
```

### Build and Clean
```bash
make all         # Run both lint and test
make clean       # Remove ./bin directory
```

## Architecture Overview

This is a Go utility library (`github.com/kaptinlin/filter`) that provides template filter functions for string manipulation, array operations, date formatting, number formatting, and mathematical computations. The codebase is organized into functional modules:

### Core Modules
- **string.go**: String manipulation functions (trim, replace, case conversions, pluralization)
- **array.go**: Slice operations (unique, join, first/last, shuffle, aggregations)
- **date.go**: Date/time formatting and manipulation using carbon library
- **number.go**: Number formatting including byte formatting
- **math.go**: Mathematical operations (abs, round, basic arithmetic)
- **data.go**: Data extraction from complex nested structures using dot notation
- **utils.go**: Common utility functions and type conversions
- **errors.go**: Unified error definitions for consistent error handling

### Key Dependencies
- `github.com/dromara/carbon/v2` - Date/time handling
- `github.com/kaptinlin/jsonpointer` - JSON pointer traversal for data extraction
- `github.com/gosimple/slug` - URL slug generation
- `github.com/jinzhu/inflection` - Word pluralization
- `github.com/dustin/go-humanize` - Human-readable formatting

### Error Handling Pattern
The codebase uses custom error types defined in errors.go:
- `ErrNotNumeric` - For non-numeric input to numeric functions
- `ErrInvalidTimeFormat` - For invalid time parsing
- `ErrUnsupportedType` - For unsupported data types
- `ErrNotSlice` - For non-slice input to slice functions
- `ErrEmptySlice` - For operations on empty slices
- `ErrInvalidArguments` - For invalid function arguments
- `ErrKeyNotFound` - For missing keys in data extraction
- `ErrIndexOutOfRange` - For invalid array/slice indices
- `ErrInvalidKeyType` - For invalid key types in data extraction
- `ErrDivisionByZero` - For division by zero operations
- `ErrModulusByZero` - For modulus by zero operations
- `ErrUnsupportedSizeType` - For unsupported types in Size filter
- `ErrNegativeValue` - For non-negative value requirements

### Testing Strategy
- Each module has corresponding `*_test.go` files with comprehensive test coverage
- `example_test.go` provides runnable Godoc examples for key functions
- Tests use `github.com/stretchr/testify` for assertions
- Race condition testing is enabled by default

### Code Quality
- Uses golangci-lint v2.9.0 with extensive linter rules enabled
- Enforces consistent code formatting with gofmt and goimports
- Custom linter configuration in `.golangci.yml` excludes certain rules for test files