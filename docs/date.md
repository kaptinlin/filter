# Date Functions in the `filter` Package

The `filter` package offers a suite of functions to work with dates in Go, making it easier to format, extract, and compute time differences.

String inputs are parsed in `time.UTC`. To format in a different zone, pass a `time.Time` already in that zone.

## Functions

### Date

Formats a timestamp into a specified format. If no format is provided, it returns a default datetime string.

#### Format Tokens

Letters not listed below pass through as literals; a backslash escapes the next byte.

| Token | Meaning                                  | Example       |
| ----- | ---------------------------------------- | ------------- |
| `Y`   | Year, 4 digits                           | `2024`        |
| `y`   | Year, 2 digits                           | `24`          |
| `m`   | Month, 2 digits                          | `03`          |
| `n`   | Month, no leading zero                   | `3`           |
| `M`   | Month, 3-letter English abbreviation     | `Mar`         |
| `F`   | Month, full English name                 | `March`       |
| `d`   | Day of month, 2 digits                   | `05`          |
| `j`   | Day of month, no leading zero            | `5`           |
| `D`   | Weekday, 3-letter English                | `Sat`         |
| `l`   | Weekday, full English name               | `Saturday`    |
| `N`   | ISO weekday (1 = Mon ... 7 = Sun)        | `6`           |
| `S`   | English ordinal suffix for day of month  | `th`          |
| `w`   | Weekday (0 = Sun ... 6 = Sat)            | `6`           |
| `z`   | Day of year, zero-based                  | `89`          |
| `W`   | ISO week of year, 2 digits               | `13`          |
| `o`   | ISO week-numbering year                  | `2024`        |
| `t`   | Days in month                            | `31`          |
| `L`   | Leap year flag                           | `1`           |
| `H`   | Hour, 24-hour, 2 digits                  | `15`          |
| `G`   | Hour, 24-hour, no leading zero           | `15`          |
| `h`   | Hour, 12-hour, 2 digits                  | `03`          |
| `g`   | Hour, 12-hour, no leading zero           | `3`           |
| `i`   | Minute, 2 digits                         | `04`          |
| `s`   | Second, 2 digits                         | `05`          |
| `u`   | Microseconds, 6 digits                   | `123456`      |
| `v`   | Milliseconds, 3 digits                   | `123`         |
| `A`   | Uppercase AM/PM                          | `PM`          |
| `a`   | Lowercase am/pm                          | `pm`          |
| `U`   | Unix timestamp, seconds                  | `1711811045`  |
| `O`   | Timezone offset                          | `+0800`       |
| `P`   | Timezone offset with colon               | `+08:00`      |
| `p`   | `P`, but UTC as `Z`                      | `Z`           |
| `T`   | Timezone abbreviation                    | `UTC`         |
| `e`   | Timezone identifier                      | `UTC`         |
| `I`   | Daylight-saving flag                     | `0`           |
| `Z`   | Timezone offset in seconds               | `28800`       |
| `c`   | ISO 8601 date                            | `2024-03-30T15:04:05+00:00` |
| `r`   | RFC 2822 date                            | `Sat, 30 Mar 2024 15:04:05 +0000` |
| `B`   | Swatch Internet time                     | `669`         |

**Example:**

```go
formatted, err := filter.Date(time.Now(), "Y-m-d")
if err != nil {
    log.Fatal(err)
}
fmt.Println(formatted) // Outputs: 2024-03-30
```

### Day

Extracts and returns the day of the month.

**Example:**

```go
day, err := filter.Day("2024-03-30")
if err != nil {
    log.Fatal(err)
}
fmt.Println(day) // Outputs: 30
```

### Month

Extracts and returns the month number.

**Example:**

```go
month, err := filter.Month(time.Now())
if err != nil {
    log.Fatal(err)
}
fmt.Println(month) // Outputs: 3 (for March)
```

### MonthFull

Returns the full month name.

**Example:**

```go
monthName, err := filter.MonthFull("2024-03-30")
if err != nil {
    log.Fatal(err)
}
fmt.Println(monthName) // Outputs: March
```

### Year

Extracts and returns the year.

**Example:**

```go
year, err := filter.Year(time.Now())
if err != nil {
    log.Fatal(err)
}
fmt.Println(year) // Outputs: 2024
```

### Week

Returns the ISO week number.

**Example:**

```go
week, err := filter.Week("2024-03-30")
if err != nil {
    log.Fatal(err)
}
fmt.Println(week) // Example output: 13
```

### Weekday

Returns the day of the week.

**Example:**

```go
weekday, err := filter.Weekday("2024-03-30")
if err != nil {
    log.Fatal(err)
}
fmt.Println(weekday) // Outputs: Saturday
```

### TimeAgo

Returns a human-readable string representing the past or future time difference between the current wall time and the input date. Past inputs render `... ago`; future inputs render `in ...`. Use `filter.TimeAgoWithClock(filter.FixedClock{T: ...}, input)` when tests need a deterministic reference point. A nil clock passed to `TimeAgoWithClock` returns an error.

**Example:**

```go
timeAgo, err := filter.TimeAgo("2024-03-01")
if err != nil {
    log.Fatal(err)
}
fmt.Println(timeAgo) // Outputs: "4 weeks ago", depending on the current date
```
