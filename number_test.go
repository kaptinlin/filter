package filter

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumber(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		format    string
		expected  string
		expectErr bool
	}{
		{
			name:     "default formatting",
			input:    12345.6789,
			format:   "",
			expected: "12,345.68",
		},
		{
			name:     "custom formatting with two decimal places",
			input:    12345.6789,
			format:   "#,###.##",
			expected: "12,345.68",
		},
		{
			name:     "no decimal places",
			input:    12345.6789,
			format:   "#,###.",
			expected: "12,346",
		},
		{
			name:     "negative number",
			input:    -12345.6789,
			format:   "#,###.##",
			expected: "-12,345.68",
		},
		{
			name:     "NaN",
			input:    math.NaN(),
			format:   "",
			expected: "NaN",
		},
		{
			name:     "Infinity",
			input:    math.Inf(1),
			format:   "",
			expected: "Infinity",
		},
		{
			name:     "-Infinity",
			input:    math.Inf(-1),
			format:   "",
			expected: "-Infinity",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Number(tt.input, tt.format)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got)
			}
		})
	}
}

func TestBytes(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expected  string
		expectErr bool
	}{
		{
			name:     "Simple number",
			input:    42,
			expected: "42 B",
		},
		{
			name:     "Kilobytes",
			input:    1024,
			expected: "1.0 kB",
		},
		{
			name:     "Megabytes",
			input:    1048576,
			expected: "1.0 MB",
		},
		{
			name:     "Gigabytes",
			input:    1073741824,
			expected: "1.1 GB",
		},
		{
			name:     "Terabytes",
			input:    1099511627776,
			expected: "1.1 TB",
		},
		{
			name:     "Petabytes",
			input:    1125899906842624,
			expected: "1.1 PB",
		},
		{
			name:      "Invalid type (string)",
			input:     "not a number",
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bytes(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got)
			}
		})
	}
}

// Benchmark tests for number operations

func BenchmarkNumber(b *testing.B) {
	for b.Loop() {
		_, _ = Number(12345.6789, "#,###.##")
	}
}

func BenchmarkBytes(b *testing.B) {
	for b.Loop() {
		_, _ = Bytes(1073741824)
	}
}
