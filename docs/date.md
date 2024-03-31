# Date Functions in the `filter` Package

The `filter` package offers a suite of functions to work with dates in Go, making it easier to format, extract, and compute time differences.

## Functions

### Date

Formats a timestamp into a specified format. If no format is provided, it returns a default datetime string.

#### Format Signs

| Sign | Description                                           | Example   |
|------|-------------------------------------------------------|-----------|
| `Y`  | Four-digit year                                       | 2024      |
| `y`  | Two-digit year                                        | 23        |
| `F`  | Full month name                                       | March     |
| `m`  | Month with leading zero                               | 03        |
| `M`  | Abbreviated month name                                | Mar       |
| `n`  | Month without leading zeros                           | 3         |
| `d`  | Day of the month with leading zero                    | 05        |
| `j`  | Day of the month without leading zeros                | 5         |
| `D`  | Abbreviated weekday name                              | Tue       |
| `l`  | Full weekday name                                     | Tuesday   |
| `a`  | Lowercase ante meridiem and post meridiem             | am, pm    |
| `A`  | Uppercase Ante meridiem and Post meridiem             | AM, PM    |
| `g`  | Hour in 12-hour format without leading zeros          | 3         |
| `h`  | Hour in 12-hour format with leading zeros             | 03        |
| `G`  | Hour in 24-hour format without leading zeros          | 15        |
| `H`  | Hour in 24-hour format with leading zeros             | 15        |
| `i`  | Minute with leading zero                              | 04        |
| `s`  | Second with leading zero                              | 05        |
| `O`  | GMT offset without colon                              | +0200     |
| `P`  | GMT offset with colon                                 | +02:00    |
| `T`  | Time zone abbreviation                                | EST       |
| `W`  | ISO-8601 week number                                  | 52        |
| `N`  | ISO-8601 numeric representation of the day of the week| 1 (for Monday) |

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
fmt.Println(weekday) // Outputs: Thursday
```

### TimeAgo

Returns a human-readable string representing the time difference between the current time and the input date.

**Example:**

```go
timeAgo, err := filter.TimeAgo("2024-03-01")
if err != nil {
    log.Fatal(err)
}
fmt.Println(timeAgo) // Outputs: "4 weeks ago", depending on the current date
```