# Number Functions in the `filter` Package

The `filter` package offers intuitive functions to format numerical values
for readability and presentation.

## Functions

### Number

Formats any numeric value using a `#,###.##`-style format string. The number
of characters after the `.` controls decimal precision; a `,` anywhere in
the integer portion enables thousands separators. Without `.`, integers render
without decimals and non-integers keep their natural precision.

**Example:**

```go
formatted, err := filter.Number(1234567.89, "#,###.##")
if err != nil {
    log.Fatal(err)
}
fmt.Println(formatted) // Outputs: "1,234,567.89"

// Integer with separators
formatted, _ = filter.Number(1234567, "#,###.")
fmt.Println(formatted) // Outputs: "1,234,567"

// Plain integer, no separators
formatted, _ = filter.Number(42, "#")
fmt.Println(formatted) // Outputs: "42"
```

### Bytes

Converts a numeric value into a human-readable byte string using SI / decimal
units (KB, MB, GB, …). Inputs must be non-negative whole-number byte counts;
negative, fractional, non-finite, and overflowing inputs return an error.

**Example:**

```go
formatted, err := filter.Bytes(1024)
if err != nil {
    log.Fatal(err)
}
fmt.Println(formatted) // Outputs: "1.0 KB"
```

For binary (KiB / MiB) output, call `humanize.BinaryBytes` from
`github.com/agentable/go-humanize` directly.
