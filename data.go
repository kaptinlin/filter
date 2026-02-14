package filter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kaptinlin/jsonpointer"
)

// Extract retrieves a value from input using dot-separated key notation.
// It supports maps, slices, arrays, structs, pointers, and interfaces.
//
// The key uses dot notation for nested access:
//   - "user.name" accesses the "name" field of "user"
//   - "items.0.title" accesses the "title" field of the first item
//   - "matrix.1.0" accesses element [1][0] of a 2D array
//
// Supported input types:
//   - map[string]interface{} and similar map types
//   - []interface{} and other slice types
//   - Arrays (including multi-dimensional)
//   - Structs with json tags or exported field names
//   - Pointers to any of the above
//   - Interfaces containing any of the above
//
// Returns ErrUnsupportedType if input is nil.
// Returns ErrKeyNotFound if the key path doesn't exist.
// Returns ErrIndexOutOfRange for invalid array/slice indices.
// Returns ErrInvalidKeyType for invalid path navigation.
func Extract(input any, key string) (any, error) {
	if input == nil {
		return nil, ErrUnsupportedType
	}

	if key == "" {
		return nil, ErrKeyNotFound
	}

	// Convert dot notation to JSON Pointer path parts
	parts := strings.Split(key, ".")

	// Use jsonpointer.Get which handles all the complex cases including array bounds
	result, err := jsonpointer.Get(input, parts...)
	if err != nil {
		// Map jsonpointer errors to filter errors for consistency
		return nil, mapJSONPointerError(err, key)
	}

	return result, nil
}

// mapJSONPointerError maps jsonpointer library errors to filter package errors
func mapJSONPointerError(err error, key string) error {
	switch {
	case errors.Is(err, jsonpointer.ErrKeyNotFound):
		return fmt.Errorf("%w: map key not found '%s'", ErrKeyNotFound, key)
	case errors.Is(err, jsonpointer.ErrFieldNotFound):
		return fmt.Errorf("%w: struct field not found '%s'", ErrKeyNotFound, key)
	case errors.Is(err, jsonpointer.ErrIndexOutOfBounds):
		return fmt.Errorf("%w: index out of range '%s'", ErrIndexOutOfRange, key)
	case errors.Is(err, jsonpointer.ErrInvalidIndex):
		return fmt.Errorf("%w: invalid index in path '%s'", ErrInvalidKeyType, key)
	case errors.Is(err, jsonpointer.ErrInvalidPathStep):
		return fmt.Errorf("%w: invalid path step in '%s'", ErrInvalidKeyType, key)
	case errors.Is(err, jsonpointer.ErrNilPointer):
		return fmt.Errorf(
			"%w: cannot navigate through nil pointer in path '%s'",
			ErrInvalidKeyType, key,
		)
	case errors.Is(err, jsonpointer.ErrNotFound):
		// For "not found" errors, check if it's trying to navigate into a primitive
		if isPrimitiveNavigationError(key) {
			return fmt.Errorf("%w: cannot navigate into primitive value", ErrInvalidKeyType)
		}
		return fmt.Errorf("%w: path not found '%s'", ErrKeyNotFound, key)
	default:
		// For unknown errors, map to invalid key type
		return fmt.Errorf("%w: %v", ErrInvalidKeyType, err) //nolint:errorlint // intentionally using %v to avoid leaking internal jsonpointer errors
	}
}

// isPrimitiveNavigationError checks if the error is due to trying to navigate into a primitive value
func isPrimitiveNavigationError(key string) bool {
	// This is a heuristic: if the key has more than one part and the error is "not found",
	// it's likely trying to navigate into a primitive value
	parts := strings.Split(key, ".")
	return len(parts) > 1
}
