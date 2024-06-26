package filter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrKeyNotFound     = errors.New("key not found")
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrInvalidKeyType  = errors.New("invalid key type")
)

// Extract retrieves a nested value from a map, slice, or array using a dot-separated key.
func Extract(input interface{}, dotKey string) (interface{}, error) {
	if input == nil {
		return nil, ErrUnsupportedType
	}

	if dotKey == "" {
		return nil, ErrKeyNotFound
	}

	parts := strings.Split(dotKey, ".")
	current := input

	for i, part := range parts {
		if current == nil {
			return nil, ErrInvalidKeyType
		}

		currentVal := reflect.ValueOf(current)

		// Early termination if the current type is not compatible with deeper examination
		if i < len(parts)-1 && currentVal.Kind() != reflect.Map && currentVal.Kind() != reflect.Slice && currentVal.Kind() != reflect.Array {
			return nil, ErrInvalidKeyType
		}

		switch currentVal.Kind() {
		case reflect.Map:
			value := currentVal.MapIndex(reflect.ValueOf(part))
			if !value.IsValid() {
				return nil, fmt.Errorf("%w: key not found '%s'", ErrKeyNotFound, part)
			}
			current = value.Interface()
		case reflect.Slice, reflect.Array:
			index, err := strconv.Atoi(part)
			if err != nil || index < 0 || index >= currentVal.Len() {
				return nil, fmt.Errorf("%w: index out of range '%s'", ErrIndexOutOfRange, part)
			}
			current = currentVal.Index(index).Interface()
		default:
			// Handle unsupported types gracefully
			return nil, fmt.Errorf("%w: unsupported type '%s'", ErrInvalidKeyType, currentVal.Kind())
		}
	}

	return current, nil
}
