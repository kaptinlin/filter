package filter

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  any
		format string
		want   string
	}{
		{"int with comma", 1234567, "#,###.", "1,234,567"},
		{"float two decimals", 1234567.89, "#,###.##", "1,234,567.89"},
		{"natural integer no comma", 1000.0, "#", "1000"},
		{"float no comma", 1234.5, "#.#", "1234.5"},
		{"negative", -1234.5, "#,###.#", "-1,234.5"},
		{"string numeric", "1234", "#,###.", "1,234"},
		{"plain integer", 42, "#", "42"},
		{"rounding carries into integer", 1.999, "#.##", "2.00"},
		{"rounding regroups integer part", 999.996, "#,###.##", "1,000.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Number(tt.input, tt.format)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNumberRejectsNonNumeric(t *testing.T) {
	t.Parallel()
	_, err := Number(struct{}{}, "#,###.##")
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"bytes under 1k", 512, "512 B"},
		{"kilobyte", 1024, "1.0 KB"},
		{"megabyte", 1048576, "1.0 MB"},
		{"gigabyte", 1073741824, "1.1 GB"},
		{"string numeric", "2048", "2.0 KB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Bytes(tt.input)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBytesRejectsNegative(t *testing.T) {
	t.Parallel()
	_, err := Bytes(-1)
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestBytesRejectsNonNumeric(t *testing.T) {
	t.Parallel()
	_, err := Bytes("not a number")
	require.ErrorIs(t, err, ErrFormat)
}

func TestBytesRejectsFractional(t *testing.T) {
	t.Parallel()
	_, err := Bytes(1.5)
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestBytesRejectsNonFinite(t *testing.T) {
	t.Parallel()
	_, err := Bytes(math.NaN())
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestBytesRejectsOverflow(t *testing.T) {
	t.Parallel()
	_, err := Bytes(uint64(1) << 63)
	require.ErrorIs(t, err, ErrInvalidInput)
}

func BenchmarkNumber(b *testing.B) {
	for b.Loop() {
		_, _ = Number(1234567.89, "#,###.##")
	}
}

func BenchmarkBytes(b *testing.B) {
	for b.Loop() {
		_, _ = Bytes(1024 * 1024 * 1024)
	}
}
