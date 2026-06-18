package filter

import "math"

// Abs returns the absolute value of input.
func Abs(input any) (float64, error) {
	v, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Abs(v), nil
}

// AtLeast returns max(input, minimum).
func AtLeast(input, minimum any) (float64, error) {
	v, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	m, err := toFloat64(minimum)
	if err != nil {
		return 0, err
	}
	return max(v, m), nil
}

// AtMost returns min(input, maximum).
func AtMost(input, maximum any) (float64, error) {
	v, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	m, err := toFloat64(maximum)
	if err != nil {
		return 0, err
	}
	return min(v, m), nil
}

// Round rounds input to the given number of decimal places.
//
// decimals accepts any numeric type (int, float, or numeric string) so
// callers from template runtimes do not need to coerce first.
func Round(input, decimals any) (float64, error) {
	v, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	d, err := toFloat64(decimals)
	if err != nil {
		return 0, err
	}
	multiplier := math.Pow(10, d)
	return math.Round(v*multiplier) / multiplier, nil
}

// Floor rounds input down to the nearest whole number.
func Floor(input any) (float64, error) {
	v, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Floor(v), nil
}

// Ceil rounds input up to the nearest whole number.
func Ceil(input any) (float64, error) {
	v, err := toFloat64(input)
	if err != nil {
		return 0, err
	}
	return math.Ceil(v), nil
}

// Plus adds addend to input.
func Plus(input, addend any) (float64, error) {
	return binaryOp(input, addend, func(a, b float64) (float64, error) {
		return a + b, nil
	})
}

// Minus subtracts subtrahend from input.
func Minus(input, subtrahend any) (float64, error) {
	return binaryOp(input, subtrahend, func(a, b float64) (float64, error) {
		return a - b, nil
	})
}

// Times multiplies input by multiplier.
func Times(input, multiplier any) (float64, error) {
	return binaryOp(input, multiplier, func(a, b float64) (float64, error) {
		return a * b, nil
	})
}

// Divide divides input by divisor.
// Returns *Error{Kind: KindArithmetic} when divisor is zero.
func Divide(input, divisor any) (float64, error) {
	return binaryOp(input, divisor, func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, arithmetic("Divide", errDivisionByZero)
		}
		return a / b, nil
	})
}

// Modulo returns the remainder of input divided by modulus.
// Returns *Error{Kind: KindArithmetic} when modulus is zero.
func Modulo(input, modulus any) (float64, error) {
	return binaryOp(input, modulus, func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, arithmetic("Modulo", errModulusByZero)
		}
		return math.Mod(a, b), nil
	})
}

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
