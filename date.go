package filter

// Date formats input with format. It returns the default date-time string when format is empty.
func Date(input any, format string) (string, error) {
	c, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	if format == "" {
		return c.ToDateTimeString(), nil
	}
	return c.Format(format), nil
}

// Day returns the day of the month.
func Day(input any) (int, error) {
	c, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return c.Day(), nil
}

// Month returns the month number.
func Month(input any) (int, error) {
	c, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return c.Month(), nil
}

// MonthFull returns the full month name.
func MonthFull(input any) (string, error) {
	c, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	return c.ToMonthString(), nil
}

// Year returns the year.
func Year(input any) (int, error) {
	c, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return c.Year(), nil
}

// Week returns the ISO week number.
func Week(input any) (int, error) {
	c, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return c.WeekOfYear(), nil
}

// Weekday returns the weekday name.
func Weekday(input any) (string, error) {
	c, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	return c.ToWeekString(), nil
}

// TimeAgo returns the difference between input and now in human-readable form.
func TimeAgo(input any) (string, error) {
	c, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	return c.DiffForHumans(), nil
}
