package filter

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// toFloat64 converts numeric input to float64. Strings are parsed as decimals.
// Returns *Error{Kind:KindInvalidInput} on unsupported types and
// *Error{Kind:KindFormat} on unparseable strings.
func toFloat64(input any) (float64, error) {
	switch v := input.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, formatErr("toFloat64", err)
		}
		return f, nil
	default:
		return 0, invalidInput("toFloat64", nil)
	}
}

func toInt64Exact(op string, input any) (int64, error) {
	switch v := input.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return uint64ToInt64(op, uint64(v))
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return uint64ToInt64(op, v)
	case float32:
		return floatToInt64Exact(op, float64(v))
	case float64:
		return floatToInt64Exact(op, v)
	case string:
		return stringToInt64Exact(op, v)
	default:
		return 0, invalidInput(op, nil)
	}
}

func uint64ToInt64(op string, v uint64) (int64, error) {
	if v > math.MaxInt64 {
		return 0, invalidInput(op, fmt.Errorf("value %d overflows int64", v))
	}
	return int64(v), nil
}

func stringToInt64Exact(op, s string) (int64, error) {
	if !strings.ContainsAny(s, ".eE") {
		v, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return v, nil
		}
		if errors.Is(err, strconv.ErrRange) {
			return 0, invalidInput(op, err)
		}
		return 0, formatErr(op, err)
	}

	return decimalStringToInt64Exact(op, s)
}

func decimalStringToInt64Exact(op, s string) (int64, error) {
	v, ok := new(big.Rat).SetString(s)
	if !ok {
		return 0, formatErr(op, fmt.Errorf("invalid number %q", s))
	}
	if !v.IsInt() {
		return 0, invalidInput(op, fmt.Errorf("expected integer, got %s", s))
	}
	if !v.Num().IsInt64() {
		return 0, invalidInput(op, fmt.Errorf("value %s overflows int64", s))
	}
	return v.Num().Int64(), nil
}

func floatToInt64Exact(op string, v float64) (int64, error) {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return 0, invalidInput(op, fmt.Errorf("expected finite number"))
	}
	if math.Trunc(v) != v {
		return 0, invalidInput(op, fmt.Errorf("expected integer, got %v", v))
	}
	if v >= 9223372036854775808.0 || v < -9223372036854775808.0 {
		return 0, invalidInput(op, fmt.Errorf("value %v overflows int64", v))
	}
	return int64(v), nil
}
