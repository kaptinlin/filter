# Golang Filter Package

The `filter` package offers a rich set of utilities for Go developers, focusing on string manipulation, array and slice operations, date and time formatting, number formatting, and mathematical computations. Its goal is to simplify handling common programming tasks. Below is an outline of available features and instructions for getting started.


## Table of Contents

- [Installing](#installing)
- [Basic Usage](#basic-usage)
- [String Functions](#string-functions)
- [Array Functions](#array-functions)
- [Date Functions](#date-functions)
- [Number Functions](#number-functions)
- [Math Functions](#math-functions)
- [Data Functions](#data-functions)

---

## Installing

Install the `filter` package with ease using the following Go command:

```bash
go get github.com/kaptinlin/filter
```

## Basic Usage

Below is an example illustrating the basic usage of the `filter` package for string manipulation:

```go
package main

import (
    "fmt"
    "github.com/kaptinlin/filter"
)

func main() {
    fmt.Println(filter.Trim("  hello world  ")) // "hello world"
    fmt.Println(filter.Replace("hello world", "world", "Go")) // "hello Go"
}
```

## String Functions

[The string Functions](docs/string.md) provide a range of functions to manipulate and query strings effectively.

| Function                                                                 | Description                                                                                      |
|--------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| [`Default`](docs/string.md#default)                                      | Returns a default value if the string is empty.                                                  |
| [`Trim`](docs/string.md#trim)                                            | Removes leading and trailing whitespace from the string.                                         |
| [`Split`](docs/string.md#split)                                          | Divides a string into a slice of strings based on a specified delimiter.                         |
| [`Replace`](docs/string.md#replace)                                      | Substitutes all instances of a specified substring with another string.                          |
| [`Remove`](docs/string.md#remove)                                        | Eliminates all occurrences of a specified substring from the string.                             |
| [`Append`](docs/string.md#append)                                        | Adds characters to the end of a string.                                                          |
| [`Prepend`](docs/string.md#prepend)                                      | Adds characters to the beginning of a string.                                                    |
| [`Length`](docs/string.md#length)                                        | Returns the number of characters in a string, accounting for UTF-8 encoding.                     |
| [`Upper`](docs/string.md#upper)                                          | Converts all characters in a string to uppercase.                                                |
| [`Lower`](docs/string.md#lower)                                          | Converts all characters in a string to lowercase.                                                |
| [`Titleize`](docs/string.md#titleize)                                    | Capitalizes the first letter of each word in a string.                                           |
| [`Capitalize`](docs/string.md#capitalize)                                | Capitalizes the first letter of a string.                                                        |
| [`Camelize`](docs/string.md#camelize)                                    | Converts a string to camelCase.                                                                  |
| [`Pascalize`](docs/string.md#pascalize)                                  | Converts a string to PascalCase.                                                                 |
| [`Dasherize`](docs/string.md#dasherize)                                  | Transforms a string into a lowercased, dash-separated format.                                    |
| [`Slugify`](docs/string.md#slugify)                                      | Converts a string into a URL-friendly "slug", ensuring it is safe for use in URLs and filenames. |
| [`Pluralize`](docs/string.md#pluralize)                                  | Determines the singular or plural form of a word based on a numeric value.                       |
| [`Ordinalize`](docs/string.md#ordinalize)                                | Converts a number to its ordinal English form.                                                   |
| [`Truncate`](docs/string.md#truncate)                                    | Shortens a string to a specified length and appends "..." if it exceeds that length.              |
| [`TruncateWords`](docs/string.md#truncatewords)                          | Truncates a string to a specified number of words, appending "..." if it exceeds that limit.     |


## Array Functions

[Array functions](docs/array.md) help you work with slices, offering tools to modify, analyze, or transform slice data.


| Function                                                               | Description                                             |
|------------------------------------------------------------------------|---------------------------------------------------------|
| [`Unique`](docs/array.md#unique)                                       | Removes duplicate elements, leaving only unique ones.   |
| [`Join`](docs/array.md#join)                                           | Concatenates slice elements into a single string.       |
| [`First`](docs/array.md#first)                                         | Retrieves the first element of the slice.               |
| [`Last`](docs/array.md#last)                                           | Returns the last element of the slice.                  |
| [`Random`](docs/array.md#random)                                       | Selects a random element from the slice.                |
| [`Reverse`](docs/array.md#reverse)                                     | Reverses the order of elements in the slice.            |
| [`Shuffle`](docs/array.md#shuffle)                                     | Randomly rearranges the elements within the slice.      |
| [`Size`](docs/array.md#size)                                           | Determines the size (length) of the slice.              |
| [`Max`](docs/array.md#max)                                             | Identifies the maximum value in a numerical slice.      |
| [`Min`](docs/array.md#min)                                             | Finds the minimum value in a numerical slice.           |
| [`Sum`](docs/array.md#sum)                                             | Calculates the sum of all elements in a numerical slice.|
| [`Average`](docs/array.md#average)                                     | Computes the average value of a numerical slice.        |
| [`Map`](docs/array.md#map)                                             | Extracts a slice of values for a specified key.         |


## Date Functions

[Date functions](docs/date.md) facilitate working with dates, including formatting, parsing, and manipulation.

| Function                                                             | Description                                                                        |
|----------------------------------------------------------------------|------------------------------------------------------------------------------------|
| [`Date`](docs/date.md#date)                                          | Formats a timestamp into a specified format or returns a default datetime string. |
| [`Day`](docs/date.md#day)                                            | Extracts and returns the day of the month.                                         |
| [`Month`](docs/date.md#month)                                        | Retrieves the month number from a date.                                            |
| [`MonthFull`](docs/date.md#monthfull)                                | Returns the full month name from a date.                                           |
| [`Year`](docs/date.md#year)                                          | Extracts and returns the year from a date.                                         |
| [`Week`](docs/date.md#week)                                          | Returns the ISO week number of a date.                                             |
| [`Weekday`](docs/date.md#weekday)                                    | Determines the day of the week from a date.                                        |
| [`TimeAgo`](docs/date.md#timeago)                                    | Provides a human-readable string representing the time difference to the present.  |


### Number Functions

[Number functions](docs/number.md) allows for the formatting of numbers for presentation and readability.

| Function                                                         | Description                                                              |
|------------------------------------------------------------------|--------------------------------------------------------------------------|
| [`Number`](docs/number.md#number)                                | Formats any numeric value based on a specified format string.            |
| [`Bytes`](docs/number.md#bytes)                                  | Converts a numeric value into a human-readable format representing bytes.|

### Math Functions

[Math functions](docs/math.md) include a variety of operations for numerical computation and manipulation.

| Function                                                         | Description                                                               |
|------------------------------------------------------------------|---------------------------------------------------------------------------|
| [`Abs`](docs/math.md#abs)                                        | Calculates the absolute value of a number.                                |
| [`AtLeast`](docs/math.md#atleast)                                | Ensures a number is at least a specified minimum.                         |
| [`AtMost`](docs/math.md#atmost)                                  | Ensures a number does not exceed a specified maximum.                     |
| [`Round`](docs/math.md#round)                                    | Rounds a number to a specified number of decimal places.                  |
| [`Floor`](docs/math.md#floor)                                    | Rounds a number down to the nearest whole number.                         |
| [`Ceil`](docs/math.md#ceil)                                      | Rounds a number up to the nearest whole number.                           |
| [`Plus`](docs/math.md#plus)                                      | Adds two numbers together.                                                |
| [`Minus`](docs/math.md#minus)                                    | Subtracts one number from another.                                        |
| [`Times`](docs/math.md#times)                                    | Multiplies two numbers.                                                    |
| [`Divide`](docs/math.md#divide)                                  | Divides one number by another, with handling for division by zero.        |
| [`Modulo`](docs/math.md#modulo)                                  | Calculates the remainder of division of one number by another.            |

### Data Functions

[Data functions](docs/data.md) provide utilities for extracting and manipulating data from complex nested structures including maps, slices, arrays, structs, pointers, and interfaces.

| Function                                                       | Description                                                           |
|----------------------------------------------------------------|-----------------------------------------------------------------------|
| [`Extract`](docs/data.md#extract)                               | Retrieves a nested value from any supported data structure using a dot-separated key path. Supports maps, slices, arrays, structs, pointers, and complex nested combinations.|


## Credits

- [go-humanize](https://github.com/dustin/go-humanize)
- [flect](https://github.com/gobuffalo/flect)
- [slug](https://github.com/gosimple/slug)
- [carbon](https://github.com/golang-module/carbon/)
- [inflection](https://github.com/jinzhu/inflection)

## How to Contribute

Contributions to the `filter` package are welcome. If you'd like to contribute, please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
