package filter

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dromara/carbon/v2"
)

// toCarbon converts input to a carbon.Carbon object, handling various input types.
func toCarbon(input any) (*carbon.Carbon, error) {
	switch v := input.(type) {
	case carbon.Carbon:
		return &v, nil
	case *carbon.Carbon:
		return v, nil
	case time.Time:
		return carbon.CreateFromStdTime(v), nil
	case string:
		parsed := carbon.Parse(v)
		if parsed.Error != nil {
			return nil, fmt.Errorf("%w: %w", ErrInvalidTimeFormat, parsed.Error)
		}
		return parsed, nil
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedType, input)
	}
}

// toFloat64 converts input to a float64.
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
			return 0, fmt.Errorf("%w: %w", ErrNotNumeric, err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("%w: got %T", ErrNotNumeric, input)
	}
}
