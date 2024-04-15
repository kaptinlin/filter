package filter

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"
)

var (
	ErrNotSlice         = errors.New("expected input to be a slice")
	ErrEmptySlice       = errors.New("slice is empty")
	ErrInvalidArguments = errors.New("invalid number of arguments")
)

// Unique removes duplicate elements from a slice.
func Unique(input interface{}) ([]interface{}, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	var uniqueItems []interface{}
	for _, item := range slice {
		if !contains(uniqueItems, item) {
			uniqueItems = append(uniqueItems, item)
		}
	}
	return uniqueItems, nil
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
	if len(slice) > 0 {
		return slice[0], nil
	}

	return nil, ErrEmptySlice
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

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(slice))))
	if err != nil {
		return nil, err
	}

	randomIndex := n.Int64()
	return slice[randomIndex], nil
}

// Reverse reverses the order of elements in a slice.
func Reverse(input interface{}) ([]interface{}, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, len(slice))

	if len(slice) == 0 {
		return result, nil
	}

	copy(result, slice)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result, nil
}

// Shuffle randomly rearranges the elements of the slice.
func Shuffle(input interface{}) ([]interface{}, error) {
	slice, err := toSlice(input)
	if err != nil {
		return nil, err
	}

	// Using crypto/rand to securely shuffle the slice.
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		// Generate a random index from 0 to i.
		num, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return nil, err // handle the error from rand.Int
		}
		j := num.Int64() // convert *big.Int to int64, then to int

		// Swap the elements at indices i and j.
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice, nil
}

// Size returns the size (length) of a slice.
func Size(input interface{}) (int, error) {
	slice, err := toSlice(input)
	if err != nil {
		return 0, err
	}
	return len(slice), nil
}

// Max returns the maximum value from a slice of float64.
func Max(input interface{}) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}

	maxVal := slice[0]
	for _, val := range slice[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, nil
}

// Min returns the minimum value from a slice of float64.
func Min(input interface{}) (float64, error) {
	slice, err := toFloat64Slice(input)
	if err != nil {
		return 0, err
	}

	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}

	minVal := slice[0]
	for _, val := range slice[1:] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal, nil
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

	sum := 0.0
	for _, val := range slice {
		sum += val
	}
	average := sum / float64(len(slice))
	return average, nil
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

// contains checks if a slice contains a specific element.
func contains(slice []interface{}, item interface{}) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
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

	var result []interface{}
	for i := 0; i < valRef.Len(); i++ {
		result = append(result, valRef.Index(i).Interface())
	}
	return result, nil
}
