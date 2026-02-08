// Package filter provides template filter functions for transforming data
// in template rendering pipelines. It supports:
//   - String manipulation: case conversion, truncation, slugification
//   - Array operations: unique, sort, shuffle, aggregation
//   - Date formatting: date parsing, component extraction, relative time
//   - Number formatting: numeric formatting, byte humanization
//   - Math operations: arithmetic, rounding, clamping
//
// All functions accept interface{} inputs for maximum flexibility
// in dynamic template contexts.
package filter

// baseAcronyms lists common acronyms that should be capitalized when found in a string.
var baseAcronyms = map[string]string{
	"ID": "ID",
}
