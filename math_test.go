package filter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "positive int",
			input:       5,
			expected:    5,
			expectError: false,
		},
		{
			name:        "negative int",
			input:       -5,
			expected:    5,
			expectError: false,
		},
		{
			name:        "positive float64",
			input:       3.14,
			expected:    3.14,
			expectError: false,
		},
		{
			name:        "negative float64",
			input:       -3.14,
			expected:    3.14,
			expectError: false,
		},
		{
			name:        "zero value",
			input:       0,
			expected:    0,
			expectError: false,
		},
		{
			name:        "string input",
			input:       "not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Abs(tt.input)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual absolute values should match.")
			}
		})
	}
}

func TestAtLeast(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		min         interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "input greater than min",
			input:       10,
			min:         5,
			expected:    10,
			expectError: false,
		},
		{
			name:        "input less than min",
			input:       2,
			min:         5,
			expected:    5,
			expectError: false,
		},
		{
			name:        "input equal to min",
			input:       5,
			min:         5,
			expected:    5,
			expectError: false,
		},
		{
			name:        "negative input and min",
			input:       -3,
			min:         -5,
			expected:    -3,
			expectError: false,
		},
		{
			name:        "input less than negative min",
			input:       -10,
			min:         -5,
			expected:    -5,
			expectError: false,
		},
		{
			name:        "string input",
			input:       "not a number",
			min:         5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "string min",
			input:       5,
			min:         "not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "both inputs invalid",
			input:       "invalid",
			min:         "also invalid",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			min:         5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil min",
			input:       5,
			min:         nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AtLeast(tt.input, tt.min)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestAtMost(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		max         interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "input greater than max",
			input:       10,
			max:         5,
			expected:    5,
			expectError: false,
		},
		{
			name:        "input less than max",
			input:       2,
			max:         5,
			expected:    2,
			expectError: false,
		},
		{
			name:        "input equal to max",
			input:       5,
			max:         5,
			expected:    5,
			expectError: false,
		},
		{
			name:        "negative input and max",
			input:       -10,
			max:         -5,
			expected:    -10,
			expectError: false,
		},
		{
			name:        "input less than negative max",
			input:       -3,
			max:         -5,
			expected:    -5,
			expectError: false,
		},
		{
			name:        "string input",
			input:       "not a number",
			max:         5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "string max",
			input:       5,
			max:         "not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "both inputs invalid",
			input:       "invalid",
			max:         "also invalid",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			max:         5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil max",
			input:       5,
			max:         nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AtMost(tt.input, tt.max)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		precision   interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "round up",
			input:       3.14159,
			precision:   2,
			expected:    3.14,
			expectError: false,
		},
		{
			name:        "round down",
			input:       3.14159,
			precision:   3,
			expected:    3.142,
			expectError: false,
		},
		{
			name:        "string round up",
			input:       "3.14159",
			precision:   2,
			expected:    3.14,
			expectError: false,
		},
		{
			name:        "string round down",
			input:       "3.14159",
			precision:   3,
			expected:    3.142,
			expectError: false,
		},
		{
			name:        "negative precision",
			input:       1234.5678,
			precision:   -2,
			expected:    1200,
			expectError: false,
		},
		{
			name:        "zero precision",
			input:       123.456,
			precision:   0,
			expected:    123,
			expectError: false,
		},
		{
			name:        "large precision",
			input:       3.14159,
			precision:   10,
			expected:    3.14159,
			expectError: false,
		},
		{
			name:        "string input",
			input:       "not a number",
			precision:   2,
			expected:    0,
			expectError: true,
		},
		{
			name:        "string precision",
			input:       3.14159,
			precision:   "two",
			expected:    0,
			expectError: true,
		},
		{
			name:        "both inputs invalid",
			input:       "invalid",
			precision:   "also invalid",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			precision:   2,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil precision",
			input:       3.14159,
			precision:   nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Round(tt.input, tt.precision)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "positive float rounding down",
			input:       3.99,
			expected:    3,
			expectError: false,
		},
		{
			name:        "negative float rounding down",
			input:       -1.1,
			expected:    -2,
			expectError: false,
		},
		{
			name:        "positive string float rounding down",
			input:       "3.99",
			expected:    3,
			expectError: false,
		},
		{
			name:        "negative string float rounding down",
			input:       "-1.1",
			expected:    -2,
			expectError: false,
		},
		{
			name:        "exact whole number",
			input:       2.0,
			expected:    2,
			expectError: false,
		},
		{
			name:        "zero",
			input:       0,
			expected:    0,
			expectError: false,
		},
		{
			name:        "string input",
			input:       "not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Floor(tt.input)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "positive float rounding up",
			input:       3.01,
			expected:    4,
			expectError: false,
		},
		{
			name:        "negative float rounding up",
			input:       -1.9,
			expected:    -1,
			expectError: false,
		},
		{
			name:        "positive string float rounding up",
			input:       "3.01",
			expected:    4,
			expectError: false,
		},
		{
			name:        "negative string float rounding up",
			input:       "-1.9",
			expected:    -1,
			expectError: false,
		},
		{
			name:        "exact whole number",
			input:       2.0,
			expected:    2,
			expectError: false,
		},
		{
			name:        "zero",
			input:       0,
			expected:    0,
			expectError: false,
		},
		{
			name:        "string input",
			input:       "not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ceil(tt.input)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestPlus(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		addend      interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "integer addition",
			input:       5,
			addend:      3,
			expected:    8,
			expectError: false,
		},
		{
			name:        "floating point addition",
			input:       2.5,
			addend:      4.5,
			expected:    7.0,
			expectError: false,
		},
		{
			name:        "negative numbers",
			input:       -1,
			addend:      -1,
			expected:    -2,
			expectError: false,
		},
		{
			name:        "integer and float addition",
			input:       10,
			addend:      0.5,
			expected:    10.5,
			expectError: false,
		},
		{
			name:        "string numbers",
			input:       "3",
			addend:      "2",
			expected:    5,
			expectError: false,
		},
		{
			name:        "string and float addition",
			input:       "5.5",
			addend:      4.5,
			expected:    10.0,
			expectError: false,
		},
		{
			name:        "invalid string input",
			input:       "not a number",
			addend:      10,
			expected:    0,
			expectError: true,
		},
		{
			name:        "one invalid string addend",
			input:       5,
			addend:      "also not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			addend:      5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil addend",
			input:       5,
			addend:      nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Plus(tt.input, tt.addend)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestMinus(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		subtrahend  interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "simple subtraction",
			input:       10,
			subtrahend:  5,
			expected:    5,
			expectError: false,
		},
		{
			name:        "subtraction resulting in negative",
			input:       5,
			subtrahend:  10,
			expected:    -5,
			expectError: false,
		},
		{
			name:        "floating point subtraction",
			input:       5.5,
			subtrahend:  0.5,
			expected:    5.0,
			expectError: false,
		},
		{
			name:        "subtraction with negative numbers",
			input:       -5,
			subtrahend:  -3,
			expected:    -2,
			expectError: false,
		},
		{
			name:        "string numbers",
			input:       "8",
			subtrahend:  "3",
			expected:    5,
			expectError: false,
		},
		{
			name:        "string and number",
			input:       "10",
			subtrahend:  2.5,
			expected:    7.5,
			expectError: false,
		},
		{
			name:        "invalid string input",
			input:       "not a number",
			subtrahend:  5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid string subtrahend",
			input:       10,
			subtrahend:  "also not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			subtrahend:  5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil subtrahend",
			input:       5,
			subtrahend:  nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Minus(tt.input, tt.subtrahend)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestTimes(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		multiplier  interface{}
		expected    float64
		expectError bool
	}{
		{
			name:        "positive numbers multiplication",
			input:       5,
			multiplier:  4,
			expected:    20,
			expectError: false,
		},
		{
			name:        "positive by negative multiplication",
			input:       10,
			multiplier:  -2,
			expected:    -20,
			expectError: false,
		},
		{
			name:        "two negative numbers multiplication",
			input:       -3,
			multiplier:  -6,
			expected:    18,
			expectError: false,
		},
		{
			name:        "multiplication with zero",
			input:       0,
			multiplier:  10,
			expected:    0,
			expectError: false,
		},
		{
			name:        "floating point multiplication",
			input:       5.5,
			multiplier:  2,
			expected:    11.0,
			expectError: false,
		},
		{
			name:        "string numbers multiplication",
			input:       "3",
			multiplier:  "4",
			expected:    12,
			expectError: false,
		},
		{
			name:        "invalid string input",
			input:       "not a number",
			multiplier:  5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid string multiplier",
			input:       5,
			multiplier:  "also not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			multiplier:  5,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil multiplier",
			input:       5,
			multiplier:  nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Times(tt.input, tt.multiplier)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		divisor     interface{}
		expected    float64
		expectError bool
		errorType   error
	}{
		{
			name:        "positive division",
			input:       10,
			divisor:     2,
			expected:    5,
			expectError: false,
		},
		{
			name:        "division resulting in fraction",
			input:       5,
			divisor:     2,
			expected:    2.5,
			expectError: false,
		},
		{
			name:        "negative division",
			input:       -6,
			divisor:     3,
			expected:    -2,
			expectError: false,
		},
		{
			name:        "division by zero",
			input:       10,
			divisor:     0,
			expected:    0,
			expectError: true,
			errorType:   ErrDivisionByZero,
		},
		{
			name:        "zero divided by any number",
			input:       0,
			divisor:     5,
			expected:    0,
			expectError: false,
		},
		{
			name:        "string numbers division",
			input:       "10",
			divisor:     "2",
			expected:    5,
			expectError: false,
		},
		{
			name:        "invalid string input",
			input:       "not a number",
			divisor:     2,
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid string divisor",
			input:       10,
			divisor:     "not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			divisor:     2,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil divisor",
			input:       10,
			divisor:     nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.input, tt.divisor)
			if tt.expectError {
				require.Error(t, err)
				if tt.errorType != nil {
					require.Equal(t, tt.errorType, err, "The expected and actual error should match.")
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}

func TestModulo(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		modulus     interface{}
		expected    float64
		expectError bool
		errorType   error
	}{
		{
			name:        "positive numbers modulus",
			input:       10,
			modulus:     3,
			expected:    1, // 10 % 3 = 1
			expectError: false,
		},
		{
			name:        "negative input modulus",
			input:       -10,
			modulus:     3,
			expected:    -1, // -10 % 3 = -1
			expectError: false,
		},
		{
			name:        "negative modulus",
			input:       10,
			modulus:     -3,
			expected:    1, // 10 % -3 = 1
			expectError: false,
		},
		{
			name:        "modulus by zero",
			input:       10,
			modulus:     0,
			expected:    0,
			expectError: true,
			errorType:   ErrModulusByZero,
		},
		{
			name:        "zero modulus any number",
			input:       0,
			modulus:     10,
			expected:    0, // 0 % 10 = 0
			expectError: false,
		},
		{
			name:        "string numbers modulus",
			input:       "10",
			modulus:     "3",
			expected:    1,
			expectError: false,
		},
		{
			name:        "invalid string input",
			input:       "not a number",
			modulus:     2,
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid string modulus",
			input:       10,
			modulus:     "also not a number",
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil input",
			input:       nil,
			modulus:     2,
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil modulus",
			input:       10,
			modulus:     nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Modulo(tt.input, tt.modulus)
			if tt.expectError {
				require.Error(t, err)
				if tt.errorType != nil {
					require.True(t, errors.Is(err, tt.errorType), "The expected and actual error should match.")
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual result should match.")
			}
		})
	}
}
