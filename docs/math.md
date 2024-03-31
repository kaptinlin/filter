# Math Functions in the `filter` Package

The `filter` package includes a set of math functions designed to perform common numerical operations in Go. These functions support a wide range of inputs, including integers, floats, and strings that can be parsed into numbers. 

## Functions

### Abs

Calculates the absolute value of a given number.

**Example:**

```go
result, err := filter.Abs(-5)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 5
```

### AtLeast

Ensures that the given number is at least as large as a minimum value.

**Example:**

```go
result, err := filter.AtLeast(3, 5)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 5
```

### AtMost

Ensures that the given number is no larger than a maximum value.

**Example:**

```go
result, err := filter.AtMost(10, 8)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 8
```

### Round

Rounds the input to the specified number of decimal places.

**Example:**

```go
result, err := filter.Round(3.14159, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 3.14
```

### Floor

Rounds the input down to the nearest whole number.

**Example:**

```go
result, err := filter.Floor(2.9)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 2
```

### Ceil

Rounds the input up to the nearest whole number.

**Example:**

```go
result, err := filter.Ceil(2.1)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 3
```

### Plus

Adds two numbers.

**Example:**

```go
result, err := filter.Plus(5, 3)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 8
```

### Minus

Subtracts the second value from the first.

**Example:**

```go
result, err := filter.Minus(10, 4)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 6
```

### Times

Multiplies the first value by the second.

**Example:**

```go
result, err := filter.Times(6, 7)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 42
```

### Divide

Divides the first value by the second. Includes handling for division by zero.

**Example:**

```go
result, err := filter.Divide(10, 0)
if err != nil {
    log.Fatal(err)
    fmt.Println(err) // Outputs: division by zero
} else {
    fmt.Println(result)
}
```

### Modulo

Returns the remainder of the division of the first value by the second.

**Example:**

```go
result, err := filter.Modulo(10, 3)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 1
```
