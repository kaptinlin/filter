package filter

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"strings"
)

// Unique removes duplicate elements from a slice.
func Unique(input interface{}) ([]interface{}, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	seen := make(map[interface{}]bool, len(slice))
	result := make([]interface{}, 0, len(slice))
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result, nil
}

// Join joins the elements of a slice into a single string with a given separator.
func Join(input interface{}, separator string) (string, error) {
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
func First(input interface{}) (interface{}, error) {
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
func Index(input interface{}, index int) (interface{}, error) {
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
func Last(input interface{}) (interface{}, error) {
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
func Random(input interface{}) (interface{}, error) {
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
func Reverse(input interface{}) ([]interface{}, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result, nil
}

// Shuffle randomly rearranges the elements of the slice.
func Shuffle(input interface{}) ([]interface{}, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	// Create a copy to avoid modifying the original
	result := slices.Clone(slice)

	// Use math/rand/v2 for efficient shuffling
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result, nil
}

// Size returns the size (length) of a slice.
func Size(input interface{}) (int, error) {
	slice, err := toSlice(input)
	if err != nil {
		return 0, err
	}
	return len(slice), nil
}

// findExtreme finds the extreme value (min or max) in a slice based on the comparison function.
func findExtreme(slice []float64, isExtreme func(current, candidate float64) bool) (float64, error) {
	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}

	extremeVal := slice[0]
	for _, val := range slice[1:] {
		if isExtreme(extremeVal, val) {
			extremeVal = val
		}
	}
	return extremeVal, nil
}

// Max returns the maximum value from a slice of float64.
func Max(input interface{}) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	return findExtreme(slice, func(current, candidate float64) bool {
		return candidate > current
	})
}

// Min returns the minimum value from a slice of float64.
func Min(input interface{}) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	return findExtreme(slice, func(current, candidate float64) bool {
		return candidate < current
	})
}

// Sum calculates the sum of all elements in a slice of float64.
func Sum(input interface{}) (float64, error) {
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
func Average(input interface{}) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}

	sum, err := Sum(input)
	if err != nil {
		return 0, err
	}

	return sum / float64(len(slice)), nil
}

// Map returns a slice of values for a specified key from each map in the input slice.
// If the key does not exist in an item, the corresponding value in the result slice will be nil.
func Map(input interface{}, key string) ([]interface{}, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0, len(slice))
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

// toFloat64Slice attempts to convert an interface{} to a slice of float64.
func toFloat64Slice(input interface{}) ([]float64, error) {
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

// toSlice attempts to convert an interface{} to a slice of interface{}.
func toSlice(input interface{}) ([]interface{}, error) {
	valRef := reflect.ValueOf(input)
	if valRef.Kind() != reflect.Slice {
		return nil, ErrNotSlice
	}

	result := make([]interface{}, valRef.Len())
	for i := 0; i < valRef.Len(); i++ {
		result[i] = valRef.Index(i).Interface()
	}
	return result, nil
}
