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

[The string functions](docs/string.md) provide a range of functions to manipulate and query strings effectively.

| Function | Description |
|---|---|
| [`Default`](docs/string.md#default) | Returns a default value if input is nil, false, or empty string. |
| [`Trim`](docs/string.md#trim) | Removes leading and trailing whitespace. |
| [`TrimLeft`](docs/string.md#trimleft) | Removes leading whitespace (Liquid `lstrip`). |
| [`TrimRight`](docs/string.md#trimright) | Removes trailing whitespace (Liquid `rstrip`). |
| [`Split`](docs/string.md#split) | Divides a string into a slice based on a delimiter. |
| [`Replace`](docs/string.md#replace) | Substitutes all occurrences of a substring. |
| [`ReplaceFirst`](docs/string.md#replacefirst) | Replaces the first occurrence of a substring. |
| [`ReplaceLast`](docs/string.md#replacelast) | Replaces the last occurrence of a substring. |
| [`Remove`](docs/string.md#remove) | Eliminates all occurrences of a substring. |
| [`RemoveFirst`](docs/string.md#removefirst) | Removes the first occurrence of a substring. |
| [`RemoveLast`](docs/string.md#removelast) | Removes the last occurrence of a substring. |
| [`Append`](docs/string.md#append) | Adds characters to the end of a string. |
| [`Prepend`](docs/string.md#prepend) | Adds characters to the beginning of a string. |
| [`Length`](docs/string.md#length) | Returns the character count, accounting for UTF-8. |
| [`Upper`](docs/string.md#upper) | Converts all characters to uppercase. |
| [`Lower`](docs/string.md#lower) | Converts all characters to lowercase. |
| [`Titleize`](docs/string.md#titleize) | Capitalizes the first letter of each word. |
| [`Capitalize`](docs/string.md#capitalize) | Capitalizes the first letter, lowercases the rest. |
| [`Camelize`](docs/string.md#camelize) | Converts a string to camelCase. |
| [`Pascalize`](docs/string.md#pascalize) | Converts a string to PascalCase. |
| [`Dasherize`](docs/string.md#dasherize) | Transforms into a lowercased, dash-separated format. |
| [`Slugify`](docs/string.md#slugify) | Converts into a URL-friendly slug. |
| [`Pluralize`](docs/string.md#pluralize) | Returns singular or plural form based on count. |
| [`Ordinalize`](docs/string.md#ordinalize) | Converts a number to its ordinal English form. |
| [`Truncate`](docs/string.md#truncate) | Shortens to a length (including ellipsis), with optional custom ellipsis. |
| [`TruncateWords`](docs/string.md#truncatewords) | Truncates to a word count, with optional custom ellipsis. |
| [`Escape`](docs/string.md#escape) | HTML-escapes `<`, `>`, `&`, `"`, `'`. |
| [`EscapeOnce`](docs/string.md#escapeonce) | HTML-escapes without double-escaping existing entities. |
| [`StripHTML`](docs/string.md#striphtml) | Removes HTML tags, scripts, styles, and comments. |
| [`StripNewlines`](docs/string.md#stripnewlines) | Removes all newline characters. |
| [`Slice`](docs/string.md#slice) | Extracts a substring or sub-slice with negative offset support. |
| [`URLEncode`](docs/string.md#urlencode) | Percent-encodes a string for URLs. |
| [`URLDecode`](docs/string.md#urldecode) | Decodes a percent-encoded string. |
| [`Base64Encode`](docs/string.md#base64encode) | Encodes a string to standard Base64. |
| [`Base64Decode`](docs/string.md#base64decode) | Decodes a standard Base64 string. |


## Array Functions

[Array functions](docs/array.md) help you work with slices, offering tools to modify, analyze, or transform slice data.

| Function | Description |
|---|---|
| [`Unique`](docs/array.md#unique) | Removes duplicate elements, leaving only unique ones. |
| [`Join`](docs/array.md#join) | Concatenates slice elements into a single string. |
| [`First`](docs/array.md#first) | Retrieves the first element of the slice. |
| [`Last`](docs/array.md#last) | Returns the last element of the slice. |
| [`Index`](docs/array.md#index) | Returns the element at a specified index. |
| [`Random`](docs/array.md#random) | Selects a random element from the slice. |
| [`Reverse`](docs/array.md#reverse) | Reverses the order of elements. |
| [`Shuffle`](docs/array.md#shuffle) | Randomly rearranges the elements. |
| [`Size`](docs/array.md#size) | Determines the size of a slice, array, or map. |
| [`Max`](docs/array.md#max) | Identifies the maximum value in a numerical slice. |
| [`Min`](docs/array.md#min) | Finds the minimum value in a numerical slice. |
| [`Sum`](docs/array.md#sum) | Calculates the sum of all elements. |
| [`Average`](docs/array.md#average) | Computes the average value. |
| [`Map`](docs/array.md#map) | Extracts values for a specified key from each element. |
| [`Sort`](docs/array.md#sort) | Sorts in ascending order, optionally by key. |
| [`SortNatural`](docs/array.md#sortnatural) | Sorts case-insensitively, optionally by key. |
| [`Compact`](docs/array.md#compact) | Removes nil elements, optionally by key. |
| [`Concat`](docs/array.md#concat) | Combines two slices into one. |
| [`Where`](docs/array.md#where) | Filters keeping elements matching a property value. |
| [`Reject`](docs/array.md#reject) | Filters removing elements matching a property value. |
| [`Find`](docs/array.md#find) | Returns first element matching a property value. |
| [`FindIndex`](docs/array.md#findindex) | Returns index of first matching element (-1 if none). |
| [`Has`](docs/array.md#has) | Checks if any element matches a property criteria. |


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
- [slug](https://github.com/gosimple/slug)
- [carbon](https://github.com/dromara/carbon)
- [inflection](https://github.com/jinzhu/inflection)
- [jsonpointer](https://github.com/kaptinlin/jsonpointer)

## How to Contribute

Contributions to the `filter` package are welcome. If you'd like to contribute, please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
