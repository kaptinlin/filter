package filter

import "errors"

var (
	// ErrNotNumeric indicates the input is not a numeric type
	ErrNotNumeric = errors.New("input is not numeric")
	// ErrInvalidTimeFormat indicates the input has an invalid time format
	ErrInvalidTimeFormat = errors.New("input has an invalid time format")
	// ErrUnsupportedType indicates the input is of an unsupported type
	ErrUnsupportedType = errors.New("input is of an unsupported type")
	// ErrNotSlice indicates the expected input should be a slice
	ErrNotSlice = errors.New("expected input to be a slice")
	// ErrEmptySlice indicates the slice is empty
	ErrEmptySlice = errors.New("slice is empty")
	// ErrInvalidArguments indicates invalid number of arguments
	ErrInvalidArguments = errors.New("invalid number of arguments")
	// ErrKeyNotFound indicates the key was not found
	ErrKeyNotFound = errors.New("key not found")
	// ErrIndexOutOfRange indicates the index is out of range
	ErrIndexOutOfRange = errors.New("index out of range")
	// ErrInvalidKeyType indicates an invalid key type
	ErrInvalidKeyType = errors.New("invalid key type")
	// ErrDivisionByZero indicates division by zero
	ErrDivisionByZero = errors.New("division by zero")
	// ErrModulusByZero indicates modulus by zero
	ErrModulusByZero = errors.New("modulus by zero")
)
