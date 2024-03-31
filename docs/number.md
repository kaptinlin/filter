# Number Functions in the `filter` Package

The `filter` package offers intuitive functions to format numerical values for readability and presentation. 

## Functions

### Number

Formats any numeric value based on a specified format string. This function allows for custom formatting, including precision, thousands separator, and more, making it highly adaptable for various display needs.

**Example:**

```go
formatted, err := filter.Number(1234567.89, "#,###.##")
if err != nil {
    log.Fatal(err)
}
fmt.Println(formatted) // Outputs: "1,234,567.89"
```

**Note:** The actual implementation of `Number` in this document assumes the existence of a `humanize.FormatFloat` method that can handle custom format strings. As this method does not exist in the current version of `github.com/dustin/go-humanize`, the example serves illustrative purposes.

### Bytes

Converts a numeric value into a human-readable format representing bytes. This function simplifies the presentation of large byte sizes, automatically choosing the appropriate unit (KB, MB, GB, etc.) based on the input's magnitude.

**Example:**

```go
formatted, err := filter.Bytes(1024)
if err != nil {
    log.Fatal(err)
}
fmt.Println(formatted) // Outputs: "1.0 kB"
```

The `Bytes` function uses the `humanize.Bytes` method from the `github.com/dustin/go-humanize` package, ensuring consistency and reliability in byte size representation.