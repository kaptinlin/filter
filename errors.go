package filter

import "errors"

var (
	// ErrNotNumeric is returned when input is not numeric.
	ErrNotNumeric = errors.New("input is not numeric")
	// ErrInvalidTimeFormat is returned when input does not contain a valid time.
	ErrInvalidTimeFormat = errors.New("input has an invalid time format")
	// ErrUnsupportedType is returned when input has an unsupported type.
	ErrUnsupportedType = errors.New("input is of an unsupported type")
	// ErrNotSlice is returned when input is not a slice or array.
	ErrNotSlice = errors.New("expected input to be a slice")
	// ErrEmptySlice is returned when an operation requires at least one element.
	ErrEmptySlice = errors.New("slice is empty")
	// ErrInvalidArguments is returned when a function receives the wrong arguments.
	ErrInvalidArguments = errors.New("invalid number of arguments")
	// ErrKeyNotFound is returned when a requested key does not exist.
	ErrKeyNotFound = errors.New("key not found")
	// ErrIndexOutOfRange is returned when an index falls outside the collection.
	ErrIndexOutOfRange = errors.New("index out of range")
	// ErrInvalidKeyType is returned when a path step cannot be applied to the current value.
	ErrInvalidKeyType = errors.New("invalid key type")
	// ErrDivisionByZero is returned when dividing by zero.
	ErrDivisionByZero = errors.New("division by zero")
	// ErrModulusByZero is returned when taking a modulus by zero.
	ErrModulusByZero = errors.New("modulus by zero")
	// ErrUnsupportedSizeType is returned when Size receives a non-collection value.
	ErrUnsupportedSizeType = errors.New("size filter expects a slice, array, or map")
	// ErrNegativeValue is returned when input must be non-negative.
	ErrNegativeValue = errors.New("input must be non-negative")
)
