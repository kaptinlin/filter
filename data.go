package filter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kaptinlin/jsonpointer"
)

// Extract returns the value at key from input using dot-separated paths.
//
// Keys can traverse maps, arrays, slices, structs, pointers, and interfaces.
// It returns ErrUnsupportedType when input is nil, ErrKeyNotFound when the
// path is missing, ErrIndexOutOfRange for invalid indexes, and
// ErrInvalidKeyType when a path step cannot be applied to the current value.
func Extract(input any, key string) (any, error) {
	if input == nil {
		return nil, ErrUnsupportedType
	}

	if key == "" {
		return nil, ErrKeyNotFound
	}

	parts := strings.Split(key, ".")

	result, err := jsonpointer.Get(input, parts...)
	if err != nil {
		return nil, mapJSONPointerError(err, key)
	}

	return result, nil
}

func mapJSONPointerError(err error, key string) error {
	switch {
	case errors.Is(err, jsonpointer.ErrKeyNotFound), errors.Is(err, jsonpointer.ErrFieldNotFound):
		return fmt.Errorf("%w: %s", ErrKeyNotFound, key)
	case errors.Is(err, jsonpointer.ErrIndexOutOfBounds):
		return fmt.Errorf("%w: %s", ErrIndexOutOfRange, key)
	case errors.Is(err, jsonpointer.ErrInvalidIndex), errors.Is(err, jsonpointer.ErrInvalidPathStep):
		return fmt.Errorf("%w: %s", ErrInvalidKeyType, key)
	case errors.Is(err, jsonpointer.ErrNilPointer):
		return fmt.Errorf("%w: cannot navigate through nil pointer in %s", ErrInvalidKeyType, key)
	case errors.Is(err, jsonpointer.ErrNotFound):
		if strings.Contains(key, ".") {
			return ErrInvalidKeyType
		}
		return fmt.Errorf("%w: %s", ErrKeyNotFound, key)
	default:
		return fmt.Errorf("%w: %w", ErrInvalidKeyType, err)
	}
}
