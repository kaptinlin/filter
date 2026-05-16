package filter

import (
	"cmp"
	"errors"
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"strings"
	"unicode/utf8"
)

// Unique removes duplicate elements while preserving first-seen order.
//
// All elements must be Comparable. Non-comparable elements (slices, maps,
// functions, structs containing them) return *Error{Kind: KindInvalidInput}
// — use UniqueBy when records should be deduplicated by a property.
func Unique(input any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return []any{}, nil
	}
	for _, item := range slice {
		if item != nil && !reflect.TypeOf(item).Comparable() {
			return nil, invalidInput("Unique", fmt.Errorf("element type %T is not comparable", item))
		}
	}

	seen := make(map[any]struct{}, len(slice))
	out := make([]any, 0, len(slice))
	for _, item := range slice {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out, nil
}

// UniqueBy removes duplicate elements by the value at key, preserving
// first-seen order. Missing or unreachable keys return an error.
func UniqueBy(input any, key string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	keys := make([]any, 0, len(slice))
	out := make([]any, 0, len(slice))
	for _, item := range slice {
		v, err := Extract(item, key)
		if err != nil {
			return nil, err
		}
		if containsValue(keys, v) {
			continue
		}
		keys = append(keys, v)
		out = append(out, item)
	}
	return out, nil
}

// First returns the first element of a slice. Empty slices return
// *Error{Kind: KindNotFound}.
func First(input any) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, notFound("First", "", nil)
	}
	return slice[0], nil
}

// Last returns the last element of a slice. Empty slices return
// *Error{Kind: KindNotFound}.
func Last(input any) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, notFound("Last", "", nil)
	}
	return slice[len(slice)-1], nil
}

// Index returns the element at i. Out-of-range returns *Error{Kind:KindNotFound}.
func Index(input any, i int) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if i < 0 || i >= len(slice) {
		return nil, notFound("Index", fmt.Sprintf("%d", i), nil)
	}
	return slice[i], nil
}

// Random returns one element chosen uniformly at random from input.
// Empty input returns *Error{Kind: KindNotFound}.
//
// Random uses math/rand/v2's package-level generator. Call RandomWithRand
// when tests or callers need reproducible output.
func Random(input any) (any, error) {
	return randomWithRand("Random", rand.IntN, input)
}

// RandomWithRand returns one element chosen by r. Passing nil returns
// *Error{Kind: KindInvalidInput}.
func RandomWithRand(r *rand.Rand, input any) (any, error) {
	if r == nil {
		return nil, invalidInput("RandomWithRand", nil)
	}
	return randomWithRand("RandomWithRand", r.IntN, input)
}

func randomWithRand(op string, intN func(int) int, input any) (any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, notFound(op, "", nil)
	}
	return slice[intN(len(slice))], nil
}

// Shuffle returns a new slice with elements rearranged in random order.
// Shuffle uses math/rand/v2's package-level generator. Call ShuffleWithRand
// when tests or callers need reproducible output.
func Shuffle(input any) ([]any, error) {
	return shuffleWithRand(rand.Shuffle, input)
}

// ShuffleWithRand returns a new slice with elements rearranged by r. Passing
// nil returns *Error{Kind: KindInvalidInput}.
func ShuffleWithRand(r *rand.Rand, input any) ([]any, error) {
	if r == nil {
		return nil, invalidInput("ShuffleWithRand", nil)
	}
	return shuffleWithRand(r.Shuffle, input)
}

func shuffleWithRand(shuffle func(int, func(int, int)), input any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := slices.Clone(slice)
	shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	return out, nil
}

// Concat returns a new slice containing the elements of input followed by
// the elements of other. Both must be slices or arrays.
func Concat(input, other any) ([]any, error) {
	a, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	b, err := toSlice(other)
	if err != nil {
		return nil, err
	}
	out := slices.Concat(a, b)
	if out == nil {
		out = []any{}
	}
	return out, nil
}

// Join concatenates the elements of input into a single string with separator
// between each element. Elements are stringified with fmt.Sprint.
//
// An empty separator is allowed and equivalent to no separator.
func Join(input any, separator string) (string, error) {
	slice, err := toSlice(input)
	if err != nil {
		return "", err
	}
	parts := make([]string, len(slice))
	for i, item := range slice {
		parts[i] = fmt.Sprint(item)
	}
	return strings.Join(parts, separator), nil
}

// Reverse returns a new slice with elements in reverse order.
func Reverse(input any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := slices.Clone(slice)
	slices.Reverse(out)
	return out, nil
}

// Size returns the length of a collection (slice, array, or map). Strings
// return their UTF-8 rune count, matching Length.
func Size(input any) (int, error) {
	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len(), nil
	case reflect.String:
		return utf8.RuneCountInString(v.String()), nil
	default:
		return 0, invalidInput("Size", fmt.Errorf("expected slice, array, map, or string, got %T", input))
	}
}

// Sum returns the sum of all numeric elements. Empty slice returns (0, nil).
func Sum(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum, nil
}

// SumBy returns the sum of numeric values extracted at key from every element.
// Missing keys and non-numeric extracted values return errors.
func SumBy(input any, key string) (float64, error) {
	slice, err := toSlice(input)
	if err != nil {
		return 0, err
	}
	var sum float64
	for _, item := range slice {
		v, err := Extract(item, key)
		if err != nil {
			return 0, err
		}
		f, err := toFloat64(v)
		if err != nil {
			return 0, numericExtractError("SumBy", err)
		}
		sum += f
	}
	return sum, nil
}

// Average returns the mean of numeric elements. Empty slice returns
// *Error{Kind:KindInvalidInput} (no defined mean).
func Average(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}
	if len(slice) == 0 {
		return 0, invalidInput("Average", nil)
	}
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum / float64(len(slice)), nil
}

// Max returns the maximum numeric element.
func Max(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}
	if len(slice) == 0 {
		return 0, invalidInput("Max", nil)
	}
	return slices.Max(slice), nil
}

// Min returns the minimum numeric element.
func Min(input any) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}
	if len(slice) == 0 {
		return 0, invalidInput("Min", nil)
	}
	return slices.Min(slice), nil
}

// Map returns the values at key from each item in input. Items where the key
// is missing or unreachable contribute nil — this matches the Liquid
// `map: 'name'` semantics where any missing path is nil, not an error.
func Map(input any, key string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(slice))
	for _, item := range slice {
		v, err := Extract(item, key)
		if err != nil {
			out = append(out, nil)
			continue
		}
		out = append(out, v)
	}
	return out, nil
}

// Sort sorts the slice in ascending order. If key is provided, items are
// sorted by that property. Numeric values use numeric comparison; otherwise
// values are compared as strings.
func Sort(input any, key ...string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := slices.Clone(slice)
	slices.SortStableFunc(out, func(a, b any) int {
		return compareValues(extractOrSelf(a, b, key...))
	})
	return out, nil
}

// SortNatural sorts case-insensitively. If key is provided, sorts by that
// property case-insensitively.
func SortNatural(input any, key ...string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := slices.Clone(slice)
	slices.SortStableFunc(out, func(a, b any) int {
		return compareValuesNatural(extractOrSelf(a, b, key...))
	})
	return out, nil
}

// Compact removes nil elements. If key is provided, removes items where the
// property is nil or unreachable.
func Compact(input any, key ...string) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	hasKey := len(key) > 0 && key[0] != ""
	out := make([]any, 0, len(slice))
	for _, item := range slice {
		if hasKey {
			v, err := Extract(item, key[0])
			if err != nil || v == nil {
				continue
			}
		} else if item == nil {
			continue
		}
		out = append(out, item)
	}
	return out, nil
}

// Where returns items whose property at key matches the criterion.
//
// With value present: keeps items where Extract(item, key) equals value.
// With value omitted: keeps items where Extract(item, key) is truthy
// (nil and false are falsy; everything else is truthy).
func Where(input any, key string, value ...any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(slice))
	for _, item := range slice {
		if matchesCriteria(item, key, value...) {
			out = append(out, item)
		}
	}
	return out, nil
}

// Reject is the inverse of Where: returns items that do not match.
// Same overload semantics as Where.
func Reject(input any, key string, value ...any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(slice))
	for _, item := range slice {
		if !matchesCriteria(item, key, value...) {
			out = append(out, item)
		}
	}
	return out, nil
}

// Find returns the first item whose property at key equals value, or
// *Error{Kind: KindNotFound} if none matches.
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
	return nil, notFound("Find", key, nil)
}

// FindIndex returns the 0-based index of the first matching item, or -1.
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

// Has reports whether any item matches the criterion. Same overload semantics
// as Where: with value present checks equality; with value omitted checks
// truthiness of the property at key.
func Has(input any, key string, value ...any) (bool, error) {
	slice, err := toSlice(input)
	if err != nil {
		return false, err
	}
	return slices.ContainsFunc(slice, func(item any) bool {
		return matchesCriteria(item, key, value...)
	}), nil
}

// matchesCriteria reports whether item's property at key matches. With value
// present this is equality; with value omitted it is truthiness (nil and
// boolean false are falsy; everything else is truthy — including empty
// strings). Unreachable keys never match.
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

// isTruthy returns false for nil and bool(false); true for everything else.
// This is the Liquid-inspired truthy definition used by predicate filters.
func isTruthy(v any) bool {
	if v == nil {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	return true
}

// extractOrSelf extracts values at key, or returns the values unchanged when
// key is empty.
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

// compareValues sorts numbers numerically, everything else as strings.
func compareValues(a, b any) int {
	return compareValuesBy(a, b, func(s string) string { return s })
}

// compareValuesNatural compares case-insensitively after numeric coercion.
func compareValuesNatural(a, b any) int {
	return compareValuesBy(a, b, strings.ToLower)
}

func compareValuesBy(a, b any, normalize func(string) string) int {
	switch {
	case a == nil && b == nil:
		return 0
	case a == nil:
		return -1
	case b == nil:
		return 1
	}
	fa, errA := toFloat64(a)
	fb, errB := toFloat64(b)
	if errA == nil && errB == nil {
		return cmp.Compare(fa, fb)
	}
	return cmp.Compare(normalize(fmt.Sprint(a)), normalize(fmt.Sprint(b)))
}

// valuesEqual checks equality with cross-type numeric coercion.
func valuesEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable() && a == b {
		return true
	}
	fa, errA := toFloat64(a)
	fb, errB := toFloat64(b)
	if errA == nil && errB == nil {
		return fa == fb
	}
	return reflect.DeepEqual(a, b)
}

func containsValue(values []any, target any) bool {
	return slices.ContainsFunc(values, func(v any) bool {
		return valuesEqual(v, target)
	})
}

func numericExtractError(op string, err error) error {
	if errors.Is(err, ErrFormat) {
		return formatErr(op, err)
	}
	return invalidInput(op, err)
}

// toFloat64Slice converts each element of input to float64. Returns nil
// (length 0) for empty input rather than rejecting it.
func toFloat64Slice(input any) ([]float64, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}
	out := make([]float64, 0, len(slice))
	for _, item := range slice {
		v, err := toFloat64(item)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

// toSlice reflects input into a []any. Inputs that are not slice or array
// return *Error{Kind:KindInvalidInput}.
func toSlice(input any) ([]any, error) {
	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return nil, invalidInput("toSlice", fmt.Errorf("expected slice or array, got %T", input))
	}
	n := v.Len()
	out := make([]any, n)
	for i := range n {
		out[i] = v.Index(i).Interface()
	}
	return out, nil
}
