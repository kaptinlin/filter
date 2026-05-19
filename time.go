package filter

import (
	"strconv"
	"strings"
	"time"

	humanize "github.com/agentable/go-humanize"
	gotime "github.com/agentable/go-time"
)

// Clock is the source of "now" for filters that depend on the current time.
//
// Production code should use SystemClock{}; tests should inject a FixedClock
// to keep results deterministic.
type Clock interface {
	Now() time.Time
}

// SystemClock returns time.Now() in UTC.
type SystemClock struct{}

// Now satisfies Clock.
func (SystemClock) Now() time.Time { return time.Now().UTC() }

// FixedClock always returns the same instant. Use it in tests.
type FixedClock struct{ T time.Time }

// Now satisfies Clock.
func (c FixedClock) Now() time.Time { return c.T }

// Date formats input using PHP/carbon-style format tokens.
// An empty format returns the canonical "2006-01-02 15:04:05" representation.
//
// Supported tokens:
//
//	Y year (2024)            y year 2-digit (24)
//	m month 2-digit (03)     n month (3)
//	M month abbr (Mar)       F month full (March)
//	d day 2-digit (05)       j day (5)
//	D weekday abbr (Sat)     l weekday full (Saturday)
//	N ISO weekday 1-7        S ordinal suffix (st/nd/rd/th)
//	w weekday 0-6 (Sun=0)    z day of year 0-365
//	W ISO week (1-53)        o ISO week-numbering year
//	t days in month          L leap year flag (0/1)
//	H hour 24h 2-digit (07)  G hour 24h (7)
//	h hour 12h 2-digit (07)  g hour 12h (7)
//	i minute (07)            s second (07)
//	u microseconds           v milliseconds
//	A AM/PM uppercase        a am/pm lowercase
//	U Unix timestamp
//	O zone offset (+0800)    P zone offset (+08:00)
//	p P but UTC as Z         T zone abbreviation (UTC)
//	e zone identifier        I daylight-saving flag (0/1)
//	Z zone offset seconds    c ISO 8601 date
//	r RFC 2822 date          B Swatch Internet time
//
// Unknown letters pass through as literals. A backslash escapes the next byte.
func Date(input any, format string) (string, error) {
	t, err := toTime(input)
	if err != nil {
		return "", err
	}
	if format == "" {
		return t.Format("2006-01-02 15:04:05"), nil
	}
	return formatTime(t, format), nil
}

// Day returns the day of the month (1-31).
func Day(input any) (int, error) {
	t, err := toTime(input)
	if err != nil {
		return 0, err
	}
	return t.Day(), nil
}

// Month returns the month number (1-12).
func Month(input any) (int, error) {
	t, err := toTime(input)
	if err != nil {
		return 0, err
	}
	return int(t.Month()), nil
}

// MonthFull returns the full English month name (January..December).
func MonthFull(input any) (string, error) {
	t, err := toTime(input)
	if err != nil {
		return "", err
	}
	return t.Month().String(), nil
}

// Year returns the 4-digit year.
func Year(input any) (int, error) {
	t, err := toTime(input)
	if err != nil {
		return 0, err
	}
	return t.Year(), nil
}

// Week returns the ISO 8601 week-of-year (1-53).
func Week(input any) (int, error) {
	t, err := toTime(input)
	if err != nil {
		return 0, err
	}
	_, w := t.ISOWeek()
	return w, nil
}

// Weekday returns the full English weekday name (Sunday..Saturday).
func Weekday(input any) (string, error) {
	t, err := toTime(input)
	if err != nil {
		return "", err
	}
	return t.Weekday().String(), nil
}

// TimeAgo returns the difference between input and the current wall time in
// human-readable form ("3 hours ago", "in 5 minutes").
func TimeAgo(input any) (string, error) {
	return TimeAgoWithClock(SystemClock{}, input)
}

// TimeAgoWithClock returns the difference between input and clock.Now() in
// human-readable form. Tests should pass FixedClock for deterministic output.
func TimeAgoWithClock(clock Clock, input any) (string, error) {
	if clock == nil {
		return "", invalidInput("TimeAgoWithClock", nil)
	}
	t, err := toTime(input)
	if err != nil {
		return "", err
	}
	return humanize.Relative(t, clock.Now()), nil
}

// toTime coerces input to time.Time (UTC by default).
//
// Accepts time.Time, gotime.Instant/DateTime/Date, int/int64 (Unix seconds),
// float64 (Unix seconds), and strings parsed by gotime.Parse.
func toTime(input any) (time.Time, error) {
	switch v := input.(type) {
	case time.Time:
		return v, nil
	case *time.Time:
		if v == nil {
			return time.Time{}, invalidInput("toTime", nil)
		}
		return *v, nil
	case gotime.Instant:
		return v.Time(), nil
	case gotime.DateTime:
		return v.ToInstant().Time(), nil
	case gotime.Date:
		return time.Date(v.Year(), v.Month(), v.Day(), 0, 0, 0, 0, time.UTC), nil
	case int:
		return time.Unix(int64(v), 0).UTC(), nil
	case int64:
		return time.Unix(v, 0).UTC(), nil
	case int32:
		return time.Unix(int64(v), 0).UTC(), nil
	case float64:
		sec, frac := splitFloat(v)
		return time.Unix(sec, frac).UTC(), nil
	case float32:
		sec, frac := splitFloat(float64(v))
		return time.Unix(sec, frac).UTC(), nil
	case string:
		return parseTimeString(v)
	default:
		return time.Time{}, invalidInput("toTime", nil)
	}
}

func splitFloat(f float64) (sec, nsec int64) {
	sec = int64(f)
	nsec = int64((f - float64(sec)) * 1e9)
	return sec, nsec
}

func parseTimeString(s string) (time.Time, error) {
	r := gotime.Parse(s, gotime.WithZone("UTC"))
	switch r.Status {
	case gotime.Resolved:
		switch r.Kind {
		case gotime.KindInstant:
			return r.Instant().Time(), nil
		case gotime.KindDateTime:
			return r.DateTime().ToInstant().Time(), nil
		case gotime.KindDate:
			d := r.Date()
			return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC), nil
		case gotime.KindTime, gotime.KindDuration, gotime.KindInterval:
		}
	case gotime.Ambiguous, gotime.Invalid:
	}
	return time.Time{}, formatErr("toTime", invalidTimeError{s: s, cause: r.Error})
}

type invalidTimeError struct {
	s     string
	cause *gotime.TimeError
}

func (e invalidTimeError) Error() string {
	if e.cause != nil {
		return "cannot parse time " + strconv.Quote(e.s) + ": " + e.cause.Error()
	}
	return "cannot parse time " + strconv.Quote(e.s)
}

// formatTime renders t with PHP/carbon-style format tokens. Unknown letters
// pass through as literals; backslash escapes the next byte.
func formatTime(t time.Time, format string) string {
	var b strings.Builder
	b.Grow(len(format) + 8)
	for i := 0; i < len(format); i++ {
		c := format[i]
		if c == '\\' && i+1 < len(format) {
			b.WriteByte(format[i+1])
			i++
			continue
		}
		appendToken(&b, t, c)
	}
	return b.String()
}

func appendToken(b *strings.Builder, t time.Time, c byte) {
	switch c {
	case 'Y':
		b.WriteString(strconv.Itoa(t.Year()))
	case 'y':
		writePad2(b, t.Year()%100)
	case 'm':
		writePad2(b, int(t.Month()))
	case 'n':
		b.WriteString(strconv.Itoa(int(t.Month())))
	case 'M':
		b.WriteString(t.Month().String()[:3])
	case 'F':
		b.WriteString(t.Month().String())
	case 'd':
		writePad2(b, t.Day())
	case 'j':
		b.WriteString(strconv.Itoa(t.Day()))
	case 'D':
		b.WriteString(t.Weekday().String()[:3])
	case 'l':
		b.WriteString(t.Weekday().String())
	case 'N':
		w := int(t.Weekday())
		if w == 0 {
			w = 7
		}
		b.WriteString(strconv.Itoa(w))
	case 'S':
		b.WriteString(ordinalSuffix(t.Day()))
	case 'w':
		b.WriteString(strconv.Itoa(int(t.Weekday())))
	case 'z':
		b.WriteString(strconv.Itoa(t.YearDay() - 1))
	case 'W':
		_, w := t.ISOWeek()
		writePad2(b, w)
	case 'o':
		y, _ := t.ISOWeek()
		b.WriteString(strconv.Itoa(y))
	case 't':
		b.WriteString(strconv.Itoa(daysInMonth(t)))
	case 'L':
		if isLeapYear(t.Year()) {
			b.WriteByte('1')
		} else {
			b.WriteByte('0')
		}
	case 'H':
		writePad2(b, t.Hour())
	case 'G':
		b.WriteString(strconv.Itoa(t.Hour()))
	case 'h':
		writePad2(b, hour12(t.Hour()))
	case 'g':
		b.WriteString(strconv.Itoa(hour12(t.Hour())))
	case 'i':
		writePad2(b, t.Minute())
	case 's':
		writePad2(b, t.Second())
	case 'A':
		if t.Hour() < 12 {
			b.WriteString("AM")
		} else {
			b.WriteString("PM")
		}
	case 'a':
		if t.Hour() < 12 {
			b.WriteString("am")
		} else {
			b.WriteString("pm")
		}
	case 'U':
		b.WriteString(strconv.FormatInt(t.Unix(), 10))
	case 'u':
		writePadN(b, t.Nanosecond()/1000, 6)
	case 'v':
		writePadN(b, t.Nanosecond()/1_000_000, 3)
	case 'O':
		writeOffset(b, t, false)
	case 'P':
		writeOffset(b, t, true)
	case 'p':
		if _, off := t.Zone(); off == 0 {
			b.WriteByte('Z')
		} else {
			writeOffset(b, t, true)
		}
	case 'T':
		name, _ := t.Zone()
		b.WriteString(name)
	case 'e':
		b.WriteString(t.Location().String())
	case 'I':
		if t.IsDST() {
			b.WriteByte('1')
		} else {
			b.WriteByte('0')
		}
	case 'Z':
		_, off := t.Zone()
		b.WriteString(strconv.Itoa(off))
	case 'c':
		b.WriteString(t.Format("2006-01-02T15:04:05-07:00"))
	case 'r':
		b.WriteString(t.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
	case 'B':
		writePadN(b, swatchInternetTime(t), 3)
	default:
		b.WriteByte(c)
	}
}

func ordinalSuffix(day int) string {
	if day%100 >= 11 && day%100 <= 13 {
		return "th"
	}
	switch day % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func daysInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func hour12(h int) int {
	h %= 12
	if h == 0 {
		return 12
	}
	return h
}

func writePad2(b *strings.Builder, n int) {
	if n < 0 {
		b.WriteByte('-')
		n = -n
	}
	if n < 10 {
		b.WriteByte('0')
	}
	b.WriteString(strconv.Itoa(n))
}

func writeOffset(b *strings.Builder, t time.Time, colon bool) {
	_, off := t.Zone()
	if off < 0 {
		b.WriteByte('-')
		off = -off
	} else {
		b.WriteByte('+')
	}
	hour := off / 3600
	minute := off % 3600 / 60
	writePad2(b, hour)
	if colon {
		b.WriteByte(':')
	}
	writePad2(b, minute)
}

func swatchInternetTime(t time.Time) int {
	bielMeanTime := t.UTC().Add(time.Hour)
	seconds := bielMeanTime.Hour()*3600 + bielMeanTime.Minute()*60 + bielMeanTime.Second()
	return seconds * 1000 / 86400
}

func writePadN(b *strings.Builder, n, width int) {
	s := strconv.Itoa(n)
	for i := len(s); i < width; i++ {
		b.WriteByte('0')
	}
	b.WriteString(s)
}
