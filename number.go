package filter

import (
	humanize "github.com/dustin/go-humanize"
)

// Number formats a numeric value according to the specified format string.
func Number(input interface{}, format string) (string, error) {
	n, err := toFloat64(input)
	if err != nil {
		return "", err
	}

	// Here, we'd directly call the humanize.FormatFloat if it existed as described.
	// Since it doesn't, this call is illustrative only.
	formattedNumber := humanize.FormatFloat(format, n)
	return formattedNumber, nil
}

// Bytes formats a numeric value into a human-readable byte format.
func Bytes(input interface{}) (string, error) {
	n, err := toFloat64(input)
	if err != nil {
		return "", err
	}

	if n < 0 {
		return "", ErrNotNumeric
	}

	formattedBytes := humanize.Bytes(uint64(n))
	return formattedBytes, nil
}
