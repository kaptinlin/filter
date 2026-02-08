package filter

import (
	"github.com/dustin/go-humanize"
)

// Number formats a numeric value according to the specified format string.
func Number(input interface{}, format string) (string, error) {
	n, err := toFloat64(input)
	if err != nil {
		return "", err
	}

	return humanize.FormatFloat(format, n), nil
}

// Bytes formats a numeric value into a human-readable byte format.
func Bytes(input interface{}) (string, error) {
	n, err := toFloat64(input)
	if err != nil {
		return "", err
	}

	if n < 0 {
		return "", ErrNegativeValue
	}

	return humanize.Bytes(uint64(n)), nil
}
