package filter

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"strings"
)

// Unique removes duplicate elements from a slice.
func Unique(input any) ([]any, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	seen := make(map[any]bool, len(slice))
	result := make([]any, 0, len(slice))
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result, nil
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
