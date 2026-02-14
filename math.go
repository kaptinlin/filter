package filter

import (
	"math"
)

// Abs returns the absolute value of the input.
func Abs(input any) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Abs(val), nil
}

// AtLeast returns the larger of input and minimum.
func AtLeast(input, minimum any) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	min, err := toFloat64(minimum)
	if err != nil {
		return 0, err
	}
	return max(val, min), nil
}

// AtMost returns the smaller of input and maximum.
func AtMost(input, maximum any) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	max, err := toFloat64(maximum)
	if err != nil {
		return 0, err
	}
	return min(val, max), nil
}

// Round rounds the input to the specified number of decimal places.
func Round(input, precision any) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	prec, err := toFloat64(precision)
	if err != nil {
		return 0, err
	}
	multiplier := math.Pow(10, prec)
	return math.Round(val*multiplier) / multiplier, nil
}

// Floor rounds the input down to the nearest whole number.
func Floor(input any) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Floor(val), nil
}

// Ceil rounds the input up to the nearest whole number.
func Ceil(input any) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Ceil(val), nil
}

// Plus adds two numbers.
func Plus(input, addend any) (float64, error) {
	a, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	b, err := toFloat64(addend)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}

// Minus subtracts the second value from the first.
func Minus(input, subtrahend any) (float64, error) {
	a, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	b, err := toFloat64(subtrahend)
	if err != nil {
		return 0, err
	}
	return a - b, nil
}

// Times multiplies the first value by the second.
func Times(input, multiplier any) (float64, error) {
	a, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	b, err := toFloat64(multiplier)
	if err != nil {
		return 0, err
	}
	return a * b, nil
}

// Divide divides the first value by the second.
func Divide(input, divisor any) (float64, error) {
	a, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	b, err := toFloat64(divisor)
	if err != nil {
		return 0, err
	}
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Modulo returns the remainder of the division of the first value by the second.
func Modulo(input, modulus any) (float64, error) {
	a, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	b, err := toFloat64(modulus)
	if err != nil {
		return 0, err
	}
	if b == 0 {
		return 0, ErrModulusByZero
	}
	return math.Mod(a, b), nil
}
