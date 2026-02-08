package filter

import (
	"math"
)

// Abs calculates the absolute value of the input.
func Abs(input interface{}) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Abs(val), nil
}

// AtLeast ensures the input is at least as large as the minimum value.
func AtLeast(input, minimum interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	minVal, err := toFloat64(minimum)
	if err != nil {
		return 0, err
	}
	return math.Max(inputVal, minVal), nil
}

// AtMost ensures the input is no larger than the maximum value.
func AtMost(input, maximum interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	maxVal, err := toFloat64(maximum)
	if err != nil {
		return 0, err
	}
	return math.Min(inputVal, maxVal), nil
}

// Round rounds the input to the specified number of decimal places.
func Round(input, precision interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	precisionVal, err := toFloat64(precision)
	if err != nil {
		return 0, err
	}
	multiplier := math.Pow(10, precisionVal)
	return math.Round(inputVal*multiplier) / multiplier, nil
}

// Floor rounds the input down to the nearest whole number.
func Floor(input interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Floor(inputVal), nil
}

// Ceil rounds the input up to the nearest whole number.
func Ceil(input interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Ceil(inputVal), nil
}

// Plus adds two numbers.
func Plus(input, addend interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	addendVal, err := toFloat64(addend)
	if err != nil {
		return 0, err
	}
	return inputVal + addendVal, nil
}

// Minus subtracts the second value from the first.
func Minus(input, subtrahend interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	subtrahendVal, err := toFloat64(subtrahend)
	if err != nil {
		return 0, err
	}
	return inputVal - subtrahendVal, nil
}

// Times multiplies the first value by the second.
func Times(input, multiplier interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	multiplierVal, err := toFloat64(multiplier)
	if err != nil {
		return 0, err
	}
	return inputVal * multiplierVal, nil
}

// Divide divides the first value by the second, including error handling for division by zero.
func Divide(input, divisor interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	divisorVal, err := toFloat64(divisor)
	if err != nil {
		return 0, err
	}
	if divisorVal == 0 {
		return 0, ErrDivisionByZero
	}
	return inputVal / divisorVal, nil
}

// Modulo returns the remainder of the division of the first value by the second.
func Modulo(input, modulus interface{}) (float64, error) {
	inputVal, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	modulusVal, err := toFloat64(modulus)
	if err != nil {
		return 0, err
	}
	if modulusVal == 0 {
		return 0, ErrModulusByZero
	}
	return math.Mod(inputVal, modulusVal), nil
}
