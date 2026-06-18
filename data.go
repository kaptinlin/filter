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

	path, err := parsePath(key)
	if err != nil {
		return nil, invalidInputAt("Extract", key, err)
	}

	result := lookupPath(input, path)
	if result.found() {
		return result.value, nil
	}
	return nil, result.err("Extract", path)
}

type path struct {
	raw   string
	steps []string
}

func parsePath(raw string) (path, error) {
	steps := make([]string, 0, strings.Count(raw, ".")+1)
	var b strings.Builder
	b.Grow(len(raw))

	for i := 0; i < len(raw); i++ {
		switch raw[i] {
		case '.':
			steps = append(steps, b.String())
			b.Reset()
		case '\\':
			if i+1 >= len(raw) {
				return path{}, fmt.Errorf("dangling escape in path %q", raw)
			}
			next := raw[i+1]
			if next != '.' && next != '\\' {
				return path{}, fmt.Errorf("unsupported escape \\%c in path %q", next, raw)
			}
			b.WriteByte(next)
			i++
		default:
			b.WriteByte(raw[i])
		}
	}
	steps = append(steps, b.String())
	return path{raw: raw, steps: steps}, nil
}

type lookupState uint8

const (
	foundState lookupState = iota + 1
	missingState
	nullState
	invalidState
)

type lookupResult struct {
	state lookupState
	value any
	cause error
}

func foundLookup(value any) lookupResult {
	return lookupResult{state: foundState, value: value}
}

func missingLookup(cause error) lookupResult {
	return lookupResult{state: missingState, cause: cause}
}

func nullLookup(cause error) lookupResult {
	return lookupResult{state: nullState, cause: cause}
}

func invalidLookup(cause error) lookupResult {
	return lookupResult{state: invalidState, cause: cause}
}

func (r lookupResult) found() bool {
	return r.state == foundState
}

func (r lookupResult) err(op string, path path) error {
	switch r.state {
	case missingState:
		return notFound(op, path.raw, r.cause)
	case nullState, invalidState:
		return invalidInputAt(op, path.raw, r.cause)
	default:
		return invalidInputAt(op, path.raw, fmt.Errorf("unknown lookup state %d", r.state))
	}
}

func lookupPath(input any, path path) lookupResult {
	current := input
	for i, step := range path.steps {
		if step == "" {
			return missingLookup(fmt.Errorf("empty path step"))
		}
		result := lookupStep(current, step, path.raw)
		if !result.found() {
			return result
		}
		current = result.value
		if current == nil && i+1 < len(path.steps) {
			return nullLookup(fmt.Errorf("nil value at %q", path.raw))
		}
	}
	return foundLookup(current)
}

func lookupStep(input any, step, path string) lookupResult {
	v := reflect.ValueOf(input)
	if !v.IsValid() {
		return nullLookup(fmt.Errorf("nil value at %q", path))
	}
	v, result := dereferenceValue(v, path)
	if !result.found() {
		return result
	}

	switch v.Kind() {
	case reflect.Map:
		return extractMapStep(v, step, path)
	case reflect.Slice, reflect.Array:
		return extractIndexStep(v, step, path)
	case reflect.Struct:
		return extractStructStep(v, step, path)
	default:
		return invalidLookup(fmt.Errorf("cannot extract %q from %s", step, v.Kind()))
	}
}

func dereferenceValue(v reflect.Value, path string) (reflect.Value, lookupResult) {
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return reflect.Value{}, nullLookup(fmt.Errorf("nil pointer at %q", path))
		}
		v = v.Elem()
	}
	return v, foundLookup(nil)
}

func extractMapStep(v reflect.Value, step, path string) lookupResult {
	key, ok := mapKeyValue(v, step)
	if !ok {
		return invalidLookup(fmt.Errorf("cannot convert %q to %s", step, v.Type().Key()))
	}
	value := v.MapIndex(key)
	if !value.IsValid() {
		return missingLookup(fmt.Errorf("map key %q not found", step))
	}
	return foundLookup(value.Interface())
}

func mapKeyValue(v reflect.Value, step string) (reflect.Value, bool) {
	keyType := v.Type().Key()
	if keyType.Kind() == reflect.Interface && keyType.NumMethod() == 0 {
		key := reflect.ValueOf(step)
		if v.MapIndex(key).IsValid() {
			return key, true
		}
		for existing := range v.Seq2() {
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

func extractIndexStep(v reflect.Value, step, path string) lookupResult {
	index, err := strconv.Atoi(step)
	if err != nil || index < 0 {
		return invalidLookup(fmt.Errorf("invalid index %q", step))
	}
	if index >= v.Len() {
		return missingLookup(fmt.Errorf("index %d out of range", index))
	}
	return foundLookup(v.Index(index).Interface())
}

func extractStructStep(v reflect.Value, step, path string) lookupResult {
	field, ok := structFieldByPathName(v, step, path)
	if !ok {
		return missingLookup(fmt.Errorf("field %q not found", step))
	}
	return foundLookup(field.Interface())
}

func structFieldByPathName(v reflect.Value, step, path string) (reflect.Value, bool) {
	t := v.Type()
	for sf := range t.Fields() {
		if !sf.IsExported() {
			continue
		}
		name, ok := fieldPathName(&sf)
		if !ok {
			continue
		}
		if name == step || sf.Name == step {
			return v.FieldByIndex(sf.Index), true
		}
	}
	for sf := range t.Fields() {
		if !sf.Anonymous || !sf.IsExported() {
			continue
		}
		field := v.FieldByIndex(sf.Index)
		field, result := dereferenceValue(field, path)
		if !result.found() || field.Kind() != reflect.Struct {
			continue
		}
		if found, ok := structFieldByPathName(field, step, path); ok {
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
