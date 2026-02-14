package filter

// Date formats a timestamp into a specified format. Returns a string representation of the date.
func Date(input any, format string) (string, error) {
	carbonTime, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	if format == "" {
		return carbonTime.ToDateTimeString(), nil
	}
	return carbonTime.Format(format), nil
}

// Day extracts and returns the day of the month from the input date.
func Day(input any) (int, error) {
	carbonTime, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return carbonTime.Day(), nil
}

// Month extracts and returns the month number from the input date.
func Month(input any) (int, error) {
	carbonTime, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return carbonTime.Month(), nil
}

// MonthFull returns the full month name from the input date.
func MonthFull(input any) (string, error) {
	carbonTime, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	return carbonTime.ToMonthString(), nil
}

// Year extracts and returns the year from the input date.
func Year(input any) (int, error) {
	carbonTime, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return carbonTime.Year(), nil
}

// Week returns the ISO week number from the input date.
func Week(input any) (int, error) {
	carbonTime, err := toCarbon(input)
	if err != nil {
		return 0, err
	}
	return carbonTime.WeekOfYear(), nil
}

// Weekday returns the day of the week from the input date.
func Weekday(input any) (string, error) {
	c, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	return c.ToWeekString(), nil
}

// TimeAgo returns a human-readable string representing the time difference
// between the current time and the input date.
func TimeAgo(input any) (string, error) {
	carbonTime, err := toCarbon(input)
	if err != nil {
		return "", err
	}
	return carbonTime.DiffForHumans(), nil
}
