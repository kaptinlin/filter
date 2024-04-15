package filter

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-module/carbon/v2"
)

var (
	ErrNotNumeric        = errors.New("input is not numeric")
	ErrInvalidTimeFormat = errors.New("input has an invalid time format")
	ErrUnsupportedType   = errors.New("input is of an unsupported type")
)

// toCarbon converts an input of type interface{} to a carbon.Carbon object, handling various input types.
func toCarbon(input interface{}) (carbon.Carbon, error) {
	switch v := input.(type) {
	case carbon.Carbon:
		return v, nil
	case time.Time:
		return carbon.CreateFromStdTime(v), nil
	case string:
		parsedTime := carbon.Parse(v)
		if parsedTime.Error != nil {
			return carbon.Carbon{}, fmt.Errorf("%w: %s", ErrInvalidTimeFormat, parsedTime.Error) //nolint: errorlint
		}
		return parsedTime, nil
	default:
		return carbon.Carbon{}, fmt.Errorf("%w: %T", ErrUnsupportedType, input)
	}
}

// toFloat64 attempts to convert an interface{} to a float64.
func toFloat64(input interface{}) (float64, error) {
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
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("%w: received %T", ErrNotNumeric, input)
	}
}
