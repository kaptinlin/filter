package filter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Extract returns the value at key from input using dot-separated paths.
//
// Keys can traverse maps, arrays, slices, structs, pointers, and interfaces.
// Failure modes:
//   - nil input → *Error{Kind: KindInvalidInput}
//   - empty key, missing path step, or out-of-range index → *Error{Kind: KindNotFound}
//   - path step incompatible with the current value → *Error{Kind: KindInvalidInput}
//
// Extract intentionally owns this tiny path language instead of outsourcing it
// to a broader query DSL. It is a core dynamic-value transformation.
func Extract(input any, key string) (any, error) {
	if input == nil {
		return nil, invalidInput("Extract", nil)
	}
	if key == "" {
		return nil, notFound("Extract", "", nil)
	}

	steps, err := parsePath(key)
	if err != nil {
		return nil, invalidInput("Extract", err)
	}

	current := input
	for _, step := range steps {
		if step == "" {
			return nil, notFound("Extract", key, fmt.Errorf("empty path step"))
		}
		next, err := extractStep(current, step, key)
		if err != nil {
			return nil, err
		}
		current = next
	}
	return current, nil
}

func parsePath(path string) ([]string, error) {
	steps := make([]string, 0, strings.Count(path, ".")+1)
	var b strings.Builder
	b.Grow(len(path))

	for i := 0; i < len(path); i++ {
		switch path[i] {
		case '.':
			steps = append(steps, b.String())
			b.Reset()
		case '\\':
			if i+1 >= len(path) {
				return nil, fmt.Errorf("dangling escape in path %q", path)
			}
			next := path[i+1]
			if next != '.' && next != '\\' {
				return nil, fmt.Errorf("unsupported escape \\%c in path %q", next, path)
			}
			b.WriteByte(next)
			i++
		default:
			b.WriteByte(path[i])
		}
	}
	steps = append(steps, b.String())
	return steps, nil
}

func extractStep(input any, step, path string) (any, error) {
	v := reflect.ValueOf(input)
	if !v.IsValid() {
		return nil, invalidInput("Extract", fmt.Errorf("nil value at %q", path))
	}
	v, err := dereferenceValue(v, path)
	if err != nil {
		return nil, err
	}

	switch v.Kind() {
	case reflect.Map:
		return extractMapStep(v, step, path)
	case reflect.Slice, reflect.Array:
		return extractIndexStep(v, step, path)
	case reflect.Struct:
		return extractStructStep(v, step, path)
	default:
		return nil, invalidInput("Extract", fmt.Errorf("cannot extract %q from %s", step, v.Kind()))
	}
}

func dereferenceValue(v reflect.Value, path string) (reflect.Value, error) {
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return reflect.Value{}, invalidInput("Extract", fmt.Errorf("nil pointer at %q", path))
		}
		v = v.Elem()
	}
	return v, nil
}

func extractMapStep(v reflect.Value, step, path string) (any, error) {
	key, ok := mapKeyValue(v, step)
	if !ok {
		return nil, invalidInput("Extract", fmt.Errorf("cannot convert %q to %s", step, v.Type().Key()))
	}
	value := v.MapIndex(key)
	if !value.IsValid() {
		return nil, notFound("Extract", path, fmt.Errorf("map key %q not found", step))
	}
	return value.Interface(), nil
}

func mapKeyValue(v reflect.Value, step string) (reflect.Value, bool) {
	keyType := v.Type().Key()
	if keyType.Kind() == reflect.Interface && keyType.NumMethod() == 0 {
		key := reflect.ValueOf(step)
		if v.MapIndex(key).IsValid() {
			return key, true
		}
		for _, existing := range v.MapKeys() {
			if fmt.Sprint(existing.Interface()) == step {
				return existing, true
			}
		}
		return key, true
	}
	key, err := convertStringKey(step, keyType)
	if err != nil {
		return reflect.Value{}, false
	}
	return key, true
}

func convertStringKey(step string, typ reflect.Type) (reflect.Value, error) {
	switch typ.Kind() {
	case reflect.String:
		return reflect.ValueOf(step).Convert(typ), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(step, 10, typ.Bits())
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(n).Convert(typ), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n, err := strconv.ParseUint(step, 10, typ.Bits())
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(n).Convert(typ), nil
	case reflect.Bool:
		b, err := strconv.ParseBool(step)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(b).Convert(typ), nil
	default:
		return reflect.Value{}, fmt.Errorf("unsupported map key type %s", typ)
	}
}

func extractIndexStep(v reflect.Value, step, path string) (any, error) {
	index, err := strconv.Atoi(step)
	if err != nil || index < 0 {
		return nil, invalidInput("Extract", fmt.Errorf("invalid index %q", step))
	}
	if index >= v.Len() {
		return nil, notFound("Extract", path, fmt.Errorf("index %d out of range", index))
	}
	return v.Index(index).Interface(), nil
}

func extractStructStep(v reflect.Value, step, path string) (any, error) {
	field, ok := structFieldByPathName(v, step)
	if !ok {
		return nil, notFound("Extract", path, fmt.Errorf("field %q not found", step))
	}
	return field.Interface(), nil
}

func structFieldByPathName(v reflect.Value, step string) (reflect.Value, bool) {
	t := v.Type()
	for i := range t.NumField() {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}
		name, ok := fieldPathName(&sf)
		if !ok {
			continue
		}
		if name == step || sf.Name == step {
			return v.Field(i), true
		}
	}
	for i := range t.NumField() {
		sf := t.Field(i)
		if !sf.Anonymous || !sf.IsExported() {
			continue
		}
		field := v.Field(i)
		field, err := dereferenceValue(field, step)
		if err != nil || field.Kind() != reflect.Struct {
			continue
		}
		if found, ok := structFieldByPathName(field, step); ok {
			return found, true
		}
	}
	return reflect.Value{}, false
}

func fieldPathName(sf *reflect.StructField) (string, bool) {
	tag := sf.Tag.Get("json")
	if tag == "-" {
		return "", false
	}
	if tag != "" {
		name, _, _ := strings.Cut(tag, ",")
		if name != "" {
			return name, true
		}
	}
	return sf.Name, true
}
