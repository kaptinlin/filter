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

		for currentVal.Kind() == reflect.Ptr {
			if currentVal.IsNil() {
				return nil, fmt.Errorf("%w: nil pointer encountered", ErrInvalidKeyType)
			}
			currentVal = currentVal.Elem()
		}

		// Early termination if the current type is not compatible with deeper examination
		if i < len(parts)-1 && currentVal.Kind() != reflect.Map && currentVal.Kind() != reflect.Slice &&
			currentVal.Kind() != reflect.Array && currentVal.Kind() != reflect.Struct {
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
		case reflect.Struct:
			field := findFieldByJSONTag(currentVal, part)
			if !field.IsValid() {
				return nil, fmt.Errorf("%w: JSON field not found '%s'", ErrKeyNotFound, part)
			}
			if !field.CanInterface() {
				return nil, fmt.Errorf("%w: cannot access unexported field '%s'", ErrInvalidKeyType, part)
			}
			current = field.Interface()
		default:
			// Handle unsupported types gracefully
			return nil, fmt.Errorf("%w: unsupported type '%s'", ErrInvalidKeyType, currentVal.Kind())
		}
	}

	return current, nil
}

func findFieldByJSONTag(structValue reflect.Value, jsonTag string) reflect.Value {
	typ := structValue.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")

		tagParts := strings.Split(tag, ",")
		if tagParts[0] == jsonTag {
			return structValue.Field(i)
		}
	}
	return reflect.Value{}
}
