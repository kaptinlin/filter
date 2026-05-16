package filter

import (
	"math"
	"strconv"
	"strings"

	humanize "github.com/agentable/go-humanize"
)

// Number formats input according to a `#,###.##`-style format string.
//
// The format mirrors the convention historically used by `humanize.FormatFloat`
// in github.com/dustin/go-humanize so that templates such as
// `{{ value | number: '#,###.##' }}` keep producing identical output:
//
//   - The number of `#` (or any character) after the `.` controls the
//     decimal precision. Without `.`, integers render without decimals and
//     non-integers keep their natural precision.
//   - A `,` anywhere in the integer part inserts thousands separators.
//
// Examples:
//
//	Number(1234567.89, "#,###.##") → "1,234,567.89"
//	Number(1234567,    "#,###.")    → "1,234,567"
//	Number(1234.5,     "#.#")        → "1234.5"
//	Number(1234,       "#")          → "1234"
//
// Returns *Error{Kind: KindInvalidInput} for non-numeric input and
// *Error{Kind: KindFormat} for unparseable numeric strings.
func Number(input any, format string) (string, error) {
	v, err := toFloat64(input)
	if err != nil {
		return "", err
	}
	return formatNumber(v, format), nil
}

// Bytes formats a non-negative whole-number byte count using SI/decimal units
// (e.g. 1024 → "1.0 KB", 1048576 → "1.0 MB"). Negative, fractional,
// non-finite, and overflowing inputs return *Error{Kind: KindInvalidInput}.
//
// For binary (KiB / MiB) units, callers can compose humanize.BinaryBytes
// themselves; the SI form is what Liquid templates expect.
func Bytes(input any) (string, error) {
	v, err := toInt64Exact("Bytes", input)
	if err != nil {
		return "", err
	}
	if v < 0 {
		return "", invalidInput("Bytes", nil)
	}
	return formatDecimalBytes(v), nil
}

func formatDecimalBytes(v int64) string {
	s := humanize.Bytes(v)
	number, unit, ok := strings.Cut(s, " ")
	if !ok || unit == "B" || strings.Contains(number, ".") {
		return s
	}
	parsed, err := strconv.ParseFloat(number, 64)
	if err != nil || math.Abs(parsed) >= 10 {
		return s
	}
	return number + ".0 " + unit
}

// formatNumber renders v using the `#,###.##` mini-DSL. Format characters
// other than `,` and `.` are treated as placeholders — only their count after
// `.` matters (precision). Inspired by humanize.FormatFloat.
func formatNumber(v float64, format string) string {
	precision := -1
	useComma := false
	if i := strings.LastIndex(format, "."); i >= 0 {
		precision = len(format) - i - 1
		useComma = strings.Contains(format[:i], ",")
	} else {
		useComma = strings.Contains(format, ",")
	}

	if precision < 0 {
		// No decimal mark in format; render as natural number.
		if v == math.Trunc(v) && !math.IsInf(v, 0) && !math.IsNaN(v) {
			return formatFloat(v, 0, useComma)
		}
		return formatFloat(v, -1, useComma)
	}
	return formatFloat(v, precision, useComma)
}

func formatFloat(v float64, precision int, useComma bool) string {
	if math.IsNaN(v) {
		return "NaN"
	}
	if math.IsInf(v, 1) {
		return "+Inf"
	}
	if math.IsInf(v, -1) {
		return "-Inf"
	}
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	if v == 0 {
		v = 0
	}

	rendered := strconv.FormatFloat(v, 'f', precision, 64)
	if !useComma {
		return sign + rendered
	}

	intStr, fracStr, hasFrac := strings.Cut(rendered, ".")
	intStr = groupIntegerString(intStr)
	if hasFrac {
		return sign + intStr + "." + fracStr
	}
	return sign + intStr
}

func groupIntegerString(s string) string {
	if len(s) <= 3 {
		return s
	}

	first := len(s) % 3
	if first == 0 {
		first = 3
	}

	var b strings.Builder
	b.Grow(len(s) + (len(s)-1)/3)
	b.WriteString(s[:first])
	for i := first; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
