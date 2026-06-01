package filter

import (
	"errors"
	"math"
	"testing"
	"time"

	gotime "github.com/agentable/go-time"
	"github.com/stretchr/testify/require"
)

var fixedDate = time.Date(2024, time.March, 30, 15, 4, 5, 0, time.UTC)

func TestDateFormatTokens(t *testing.T) {
	t.Parallel()

	tests := []struct {
		format string
		want   string
	}{
		{"", "2024-03-30 15:04:05"},
		{"Y-m-d", "2024-03-30"},
		{"y-n-j", "24-3-30"},
		{"F j, Y", "March 30, 2024"},
		{"l, F", "Saturday, March"},
		{"D, M j", "Sat, Mar 30"},
		{"jS", "30th"},
		{"z, t, L, o", "89, 31, 1, 2024"},
		{"H:i:s", "15:04:05"},
		{"g:i A", "3:04 PM"},
		{"h:i a", "03:04 pm"},
		{"W, N", "13, 6"},
		{"w", "6"},
		{`Y\Y`, "2024Y"},
		{"U", "1711811045"},
		{"O", "+0000"},
		{"P", "+00:00"},
		{"p", "Z"},
		{"T", "UTC"},
		{"e", "UTC"},
		{"I", "0"},
		{"Z", "0"},
		{"c", "2024-03-30T15:04:05+00:00"},
		{"r", "Sat, 30 Mar 2024 15:04:05 +0000"},
		{"B", "669"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			t.Parallel()
			got, err := Date(fixedDate, tt.format)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDateSubsecondTokens(t *testing.T) {
	t.Parallel()

	value := time.Date(2024, time.March, 30, 15, 4, 5, 123456789, time.UTC)
	got, err := Date(value, "u v")
	require.NoError(t, err)
	require.Equal(t, "123456 123", got)
}

func TestDateTimezoneTokens(t *testing.T) {
	t.Parallel()

	value := time.Date(2024, time.March, 30, 23, 4, 5, 0, time.FixedZone("CST", 8*60*60))
	got, err := Date(value, "O P p T e Z")
	require.NoError(t, err)
	require.Equal(t, "+0800 +08:00 +08:00 CST CST 28800", got)
}

func TestDateAcceptsMultipleInputTypes(t *testing.T) {
	t.Parallel()

	unix := fixedDate.Unix()
	tests := []struct {
		name  string
		input any
	}{
		{"time.Time", fixedDate},
		{"int64 unix", unix},
		{"int unix", int(unix)},
		{"string RFC3339", "2024-03-30T15:04:05Z"},
		{"string date", "2024-03-30"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Date(tt.input, "Y-m-d")
			require.NoError(t, err)
			require.Equal(t, "2024-03-30", got)
		})
	}
}

func TestDateAcceptsGoTimeTypes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
	}{
		{name: "instant", input: gotime.InstantFromTime(fixedDate)},
		{name: "date time", input: gotime.DateTimeFromTime(fixedDate, gotime.UTC)},
		{name: "date", input: gotime.DateFromTime(fixedDate)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Date(tt.input, "Y-m-d")
			require.NoError(t, err)
			require.Equal(t, "2024-03-30", got)
		})
	}
}

func TestDatePreservesGoTimeDateTimeZone(t *testing.T) {
	t.Parallel()

	zone, err := gotime.LoadZone("Asia/Shanghai")
	require.NoError(t, err)
	input := gotime.DateTimeFromTime(time.Date(2024, time.March, 30, 23, 4, 5, 0, zone.Location()), zone)

	got, err := Date(input, "Y-m-d H:i P e")
	require.NoError(t, err)
	require.Equal(t, "2024-03-30 23:04 +08:00 Asia/Shanghai", got)
}

func TestDateRejectsNonFiniteUnixSeconds(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
	}{
		{name: "nan float64", input: math.NaN()},
		{name: "positive infinity", input: math.Inf(1)},
		{name: "negative infinity", input: math.Inf(-1)},
		{name: "nan float32", input: float32(math.NaN())},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := Date(tt.input, "Y-m-d")
			require.ErrorIs(t, err, ErrInvalidInput)
		})
	}
}

func TestDateRejectsUnparseableString(t *testing.T) {
	t.Parallel()
	_, err := Date("not-a-date", "Y-m-d")
	require.ErrorIs(t, err, ErrFormat)
}

func TestDateRejectsUnsupportedType(t *testing.T) {
	t.Parallel()
	_, err := Date(struct{}{}, "Y-m-d")
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestDateComponents(t *testing.T) {
	t.Parallel()

	day, err := Day(fixedDate)
	require.NoError(t, err)
	require.Equal(t, 30, day)

	mon, err := Month(fixedDate)
	require.NoError(t, err)
	require.Equal(t, 3, mon)

	mf, err := MonthFull(fixedDate)
	require.NoError(t, err)
	require.Equal(t, "March", mf)

	y, err := Year(fixedDate)
	require.NoError(t, err)
	require.Equal(t, 2024, y)

	w, err := Week(fixedDate)
	require.NoError(t, err)
	require.Equal(t, 13, w)

	wd, err := Weekday(fixedDate)
	require.NoError(t, err)
	require.Equal(t, "Saturday", wd)
}

func TestTimeAgo(t *testing.T) {
	t.Parallel()

	ref := time.Date(2024, time.March, 30, 16, 0, 0, 0, time.UTC)
	clock := FixedClock{T: ref}
	target := time.Date(2024, time.March, 30, 15, 0, 0, 0, time.UTC)

	got, err := TimeAgoWithClock(clock, target)
	require.NoError(t, err)
	require.Equal(t, "1 hour ago", got)
}

func TestTimeAgoRejectsNilClock(t *testing.T) {
	t.Parallel()
	_, err := TimeAgoWithClock(nil, fixedDate)
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestSystemClockReturnsUTC(t *testing.T) {
	t.Parallel()
	now := SystemClock{}.Now()
	require.Equal(t, time.UTC, now.Location())
}

func TestDateUnknownTokenPassesThrough(t *testing.T) {
	t.Parallel()
	got, err := Date(fixedDate, "Y-?-d")
	require.NoError(t, err)
	require.Equal(t, "2024-?-30", got)
}

func TestDateErrorIsKindFormat(t *testing.T) {
	t.Parallel()
	_, err := Date("xyz", "Y")
	var fe *Error
	require.True(t, errors.As(err, &fe))
	require.Equal(t, KindFormat, fe.Kind)
}

func FuzzDateFormat(f *testing.F) {
	f.Add("Y-m-d H:i:s")
	f.Add(`\Y\m\d`)
	f.Add("F j, Y")
	f.Add("")
	f.Fuzz(func(t *testing.T, format string) {
		_, _ = Date(fixedDate, format)
	})
}

func BenchmarkDateFormat(b *testing.B) {
	for b.Loop() {
		_, _ = Date(fixedDate, "F j, Y g:i A")
	}
}
