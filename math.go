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
	minVal, err := toFloat64(minimum)
	if err != nil {
		return 0, err
	}
	return max(val, minVal), nil
}

// AtMost returns the smaller of input and maximum.
func AtMost(input, maximum any) (float64, error) {
	val, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	maxVal, err := toFloat64(maximum)
	if err != nil {
		return 0, err
	}
	return min(val, maxVal), nil
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

// binaryOp is a helper for binary arithmetic operations.
func binaryOp(a, b any, op func(float64, float64) (float64, error)) (float64, error) {
	x, err := toFloat64(a)
	if err != nil {
		return 0, err
	}
	y, err := toFloat64(b)
	if err != nil {
		return 0, err
	}
	return op(x, y)
}

// Plus adds two numbers.
func Plus(input, addend any) (float64, error) {
	return binaryOp(input, addend, func(a, b float64) (float64, error) {
		return a + b, nil
	})
}

// Minus subtracts the second value from the first.
func Minus(input, subtrahend any) (float64, error) {
	return binaryOp(input, subtrahend, func(a, b float64) (float64, error) {
		return a - b, nil
	})
}

// Times multiplies the first value by the second.
func Times(input, multiplier any) (float64, error) {
	return binaryOp(input, multiplier, func(a, b float64) (float64, error) {
		return a * b, nil
	})
}

// Divide divides the first value by the second.
func Divide(input, divisor any) (float64, error) {
	return binaryOp(input, divisor, func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	})
}

// Modulo returns the remainder of the division of the first value by the second.
func Modulo(input, modulus any) (float64, error) {
	return binaryOp(input, modulus, func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, ErrModulusByZero
		}
		return math.Mod(a, b), nil
	})
}
