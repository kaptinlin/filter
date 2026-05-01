package filter

import (
	"cmp"
	"encoding/binary"
	"fmt"
	"hash/maphash"
	"maps"
	"math"
	"math/rand/v2"
	"reflect"
	"slices"
	"strings"
)

// Unique removes duplicate elements while preserving their first-seen order.
// Comparable values use a map fast path; non-comparable values fall back to hashing.
func Unique(input any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	if len(slice) == 0 {
		return []any{}, nil
	}

	allComparable := true
	for _, item := range slice {
		if item != nil && !reflect.TypeOf(item).Comparable() {
			allComparable = false
			break
		}
	}

	if allComparable {
		return uniqueComparable(slice), nil
	}
	return uniqueHash(slice), nil
}

func uniqueComparable(slice []any) []any {
	seen := make(map[any]struct{}, len(slice))
	result := make([]any, 0, len(slice))
	for _, item := range slice {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func uniqueHash(slice []any) []any {
	hashes := make(map[uint64][]int, len(slice))
	seed := maphash.MakeSeed()
	result := make([]any, 0, len(slice))

	for _, item := range slice {
		var h maphash.Hash
		h.SetSeed(seed)
		hashValue(&h, item)
		hv := h.Sum64()

		duplicate := false
		if indices := hashes[hv]; len(indices) > 0 {
			for _, j := range indices {
				if deepEqualValue(item, result[j]) {
					duplicate = true
					break
				}
			}
		}

		if !duplicate {
			hashes[hv] = append(hashes[hv], len(result))
			result = append(result, item)
		}
	}
	return result
}

func hashValue(h *maphash.Hash, v any) {
	switch val := v.(type) {
	case nil:
		_ = h.WriteByte(0)

	case bool:
		if val {
			_ = h.WriteByte(1)
		} else {
			_ = h.WriteByte(0)
		}

	case float64:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], math.Float64bits(val))
		_, _ = h.Write(buf[:])

	case int:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(val)) //nolint:gosec // Overflow is acceptable for hashing
		_, _ = h.Write(buf[:])

	case string:
		_, _ = h.WriteString(val)

	case []any:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(len(val)))
		_, _ = h.Write(buf[:])
		for _, item := range val {
			hashValue(h, item)
		}

	case map[string]any:
		keys := slices.Sorted(maps.Keys(val))

		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(len(keys)))
		_, _ = h.Write(buf[:])

		for _, k := range keys {
			_, _ = h.WriteString(k)
			hashValue(h, val[k])
		}

	default:
		hashValueReflect(h, reflect.ValueOf(v))
	}
}

func hashValueReflect(h *maphash.Hash, rv reflect.Value) {
	if !rv.IsValid() {
		_ = h.WriteByte(0)
		return
	}

	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			_ = h.WriteByte(1)
		} else {
			_ = h.WriteByte(0)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(rv.Int())) //nolint:gosec // Overflow is acceptable for hashing
		_, _ = h.Write(buf[:])

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], rv.Uint())
		_, _ = h.Write(buf[:])

	case reflect.Float32, reflect.Float64:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], math.Float64bits(rv.Float()))
		_, _ = h.Write(buf[:])

	case reflect.String:
		_, _ = h.WriteString(rv.String())

	case reflect.Slice, reflect.Array:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(rv.Len())) //nolint:gosec // Overflow is acceptable for hashing
		_, _ = h.Write(buf[:])
		for i := range rv.Len() {
			hashValueReflect(h, rv.Index(i))
		}

	case reflect.Map:
		keys := rv.MapKeys()
		slices.SortFunc(keys, func(a, b reflect.Value) int {
			return strings.Compare(fmt.Sprint(a.Interface()), fmt.Sprint(b.Interface()))
		})

		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(len(keys)))
		_, _ = h.Write(buf[:])

		for _, k := range keys {
			hashValueReflect(h, k)
			hashValueReflect(h, rv.MapIndex(k))
		}

	case reflect.Interface, reflect.Pointer:
		if rv.IsNil() {
			_ = h.WriteByte(0)
		} else {
			hashValueReflect(h, rv.Elem())
		}

	case reflect.Invalid, reflect.Uintptr, reflect.Complex64, reflect.Complex128,
		reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer:
		_, _ = fmt.Fprint(h, rv.Interface())
	}
}

func deepEqualValue(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	switch va := a.(type) {
	case bool:
		vb, ok := b.(bool)
		return ok && va == vb

	case float64:
		vb, ok := b.(float64)
		return ok && va == vb

	case int:
		vb, ok := b.(int)
		return ok && va == vb

	case string:
		vb, ok := b.(string)
		return ok && va == vb

	case []any:
		vb, ok := b.([]any)
		return ok && slices.EqualFunc(va, vb, deepEqualValue)

	case map[string]any:
		vb, ok := b.(map[string]any)
		return ok && maps.EqualFunc(va, vb, deepEqualValue)
	}

	return reflect.DeepEqual(a, b)
}

// Join joins the elements of a slice into a single string with a given separator.
func Join(input any, separator string) (string, error) {
	if separator == "" {
		return "", ErrInvalidArguments
	}

	slice, err := toSlice(input)
	if err != nil {
		return "", err
	}

	strs := make([]string, 0, len(slice))
	for _, item := range slice {
		strs = append(strs, fmt.Sprint(item))
	}

	return strings.Join(strs, separator), nil
}

// First returns the first element of a slice.
func First(input any) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, ErrEmptySlice
	}
	return slice[0], nil
}

// Index returns the element at a specified index in a slice.
func Index(input any, index int) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	if index < 0 || index >= len(slice) {
		return nil, ErrIndexOutOfRange
	}

	return slice[index], nil
}

// Last returns the last element of a slice.
func Last(input any) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, ErrEmptySlice
	}
	return slice[len(slice)-1], nil
}

// Random selects a random element from a slice.
func Random(input any) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, ErrEmptySlice
	}

	randomIndex := rand.IntN(len(slice)) //nolint:gosec // math/rand/v2 is acceptable for non-cryptographic use
	return slice[randomIndex], nil
}

// Reverse reverses the order of elements in a slice.
func Reverse(input any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	result := slices.Clone(slice)
	slices.Reverse(result)
	return result, nil
}

// Shuffle randomly rearranges the elements of the slice.
func Shuffle(input any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	result := slices.Clone(slice)

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result, nil
}

// Size returns the length of a collection (slice, array, or map).
// For string length, use [Length] instead.
func Size(input any) (int, error) {
	val := reflect.ValueOf(input)
	kind := val.Kind()

	if kind == reflect.Slice || kind == reflect.Array || kind == reflect.Map {
		return val.Len(), nil
	}

	return 0, fmt.Errorf("%w: got %T", ErrUnsupportedSizeType, input)
}

// Max returns the maximum value from a slice of float64.
func Max(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}

	return slices.Max(slice), nil
}

// Min returns the minimum value from a slice of float64.
func Min(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}

	return slices.Min(slice), nil
}

// Sum calculates the sum of all elements in a slice of float64.
func Sum(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	var sum float64
	for _, val := range slice {
		sum += val
	}
	return sum, nil
}

// Average calculates the average value of elements in a slice of float64.
func Average(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}

	var sum float64
	for _, val := range slice {
		sum += val
	}
	return sum / float64(len(slice)), nil
}

// Map returns a slice of values for a specified key from each map in the input slice.
// If the key does not exist in an item, the corresponding value in the result slice will be nil.
func Map(input any, key string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	result := make([]any, 0, len(slice))
	for _, item := range slice {
		value, err := Extract(item, key)
		if err != nil {
			result = append(result, nil)
			continue
		}
		result = append(result, value)
	}
	return result, nil
}

// Sort sorts a slice in ascending order.
// If key is provided, sorts slice of maps/structs by that property.
// Elements whose key cannot be extracted retain their relative order.
func Sort(input any, key ...string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	result := slices.Clone(slice)
	slices.SortStableFunc(result, func(a, b any) int {
		return compareValues(extractOrSelf(a, b, key...))
	})
	return result, nil
}

// SortNatural sorts a slice case-insensitively.
// If key is provided, sorts by that property case-insensitively.
// Elements whose key cannot be extracted retain their relative order.
func SortNatural(input any, key ...string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	result := slices.Clone(slice)
	slices.SortStableFunc(result, func(a, b any) int {
		return compareValuesNatural(extractOrSelf(a, b, key...))
	})
	return result, nil
}

// Compact removes nil elements from a slice.
// If key is provided, removes elements where the property is nil.
func Compact(input any, key ...string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	hasKey := len(key) > 0 && key[0] != ""
	result := make([]any, 0, len(slice))
	for _, item := range slice {
		if hasKey {
			v, err := Extract(item, key[0])
			if err != nil || v == nil {
				continue
			}
		} else if item == nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

// Concat combines two slices into one.
func Concat(input, other any) ([]any, error) {
	sliceA, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	sliceB, err := toSlice(other)
	if err != nil {
		return nil, err
	}
	result := slices.Concat(sliceA, sliceB)
	if result == nil {
		result = []any{}
	}
	return result, nil
}

// Where filters a slice, keeping elements where the given property equals the given value.
// If value is omitted, keeps elements where the property is truthy.
func Where(input any, key string, value ...any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	result := make([]any, 0, len(slice))
	for _, item := range slice {
		if matchesCriteria(item, key, value...) {
			result = append(result, item)
		}
	}
	return result, nil
}

// Reject filters a slice, removing elements where the given property equals the given value.
// If value is omitted, removes elements where the property is truthy.
func Reject(input any, key string, value ...any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	result := make([]any, 0, len(slice))
	for _, item := range slice {
		if !matchesCriteria(item, key, value...) {
			result = append(result, item)
		}
	}
	return result, nil
}

// Find returns the first element in a slice where the given property equals the given value.
func Find(input any, key string, value any) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	for _, item := range slice {
		v, err := Extract(item, key)
		if err != nil {
			continue
		}
		if valuesEqual(v, value) {
			return item, nil
		}
	}
	return nil, nil
}

// FindIndex returns the 0-based index of the first element where the given property
// equals the given value. Returns -1 if not found.
func FindIndex(input any, key string, value any) (int, error) {
	slice, err := toSlice(input)
	if err != nil {
		return -1, err
	}
	return slices.IndexFunc(slice, func(item any) bool {
		v, err := Extract(item, key)
		return err == nil && valuesEqual(v, value)
	}), nil
}

// Has returns true if any element in the slice has a property matching the given criteria.
// If value is provided, checks property == value.
// If value is omitted, checks property is truthy.
func Has(input any, key string, value ...any) (bool, error) {
	slice, err := toSlice(input)
	if err != nil {
		return false, err
	}
	return slices.ContainsFunc(slice, func(item any) bool {
		return matchesCriteria(item, key, value...)
	}), nil
}

// extractOrSelf extracts values at key, or returns the values unchanged when key is empty.
func extractOrSelf(a, b any, key ...string) (any, any) {
	if len(key) > 0 && key[0] != "" {
		va, errA := Extract(a, key[0])
		vb, errB := Extract(b, key[0])
		if errA != nil {
			va = nil
		}
		if errB != nil {
			vb = nil
		}
		return va, vb
	}
	return a, b
}

// matchesCriteria checks whether item's property at key matches the given criteria.
// If value is provided, checks property == value; otherwise checks property is truthy.
// Returns false if the key cannot be extracted.
func matchesCriteria(item any, key string, value ...any) bool {
	v, err := Extract(item, key)
	if err != nil {
		return false
	}
	if len(value) > 0 {
		return valuesEqual(v, value[0])
	}
	return isTruthy(v)
}

// compareValues compares two values for sorting.
// Numbers are compared numerically, everything else as strings.
func compareValues(a, b any) int {
	return compareValuesBy(a, b, func(s string) string { return s })
}

// compareValuesNatural compares two values case-insensitively for natural sorting.
func compareValuesNatural(a, b any) int {
	return compareValuesBy(a, b, strings.ToLower)
}

func compareValuesBy(a, b any, normalize func(string) string) int {
	if a == nil && b == nil {
		return 0
	}
	if a == nil {
		return -1
	}
	if b == nil {
		return 1
	}
	fa, errA := toFloat64(a)
	fb, errB := toFloat64(b)
	if errA == nil && errB == nil {
		return cmp.Compare(fa, fb)
	}
	return cmp.Compare(normalize(fmt.Sprint(a)), normalize(fmt.Sprint(b)))
}

// valuesEqual checks if two values are equal, handling numeric type differences.
func valuesEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)
	if ta.Comparable() && tb.Comparable() && a == b {
		return true
	}

	fa, errA := toFloat64(a)
	fb, errB := toFloat64(b)
	if errA == nil && errB == nil {
		return fa == fb
	}
	return reflect.DeepEqual(a, b)
}

// isTruthy returns true if the value is not nil and not false.
func isTruthy(v any) bool {
	if v == nil {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	return true
}

// toFloat64Slice converts input to a slice of float64.
func toFloat64Slice(input any) ([]float64, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	result := make([]float64, 0, len(slice))
	for _, item := range slice {
		val, err := toFloat64(item)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

// toSlice converts input to a slice of any.
func toSlice(input any) ([]any, error) {
	v := reflect.ValueOf(input)
	kind := v.Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrNotSlice
	}

	length := v.Len()
	result := make([]any, length)
	for i := range length {
		result[i] = v.Index(i).Interface()
	}
	return result, nil
}
