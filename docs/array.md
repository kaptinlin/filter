# Array Functions in the `filter` Package

The `filter` package provides a collection of functions designed to manipulate and analyze slices in Go effectively. 

## Functions

### Unique

Removes duplicate elements from a slice, returning a slice with only unique elements.

**Example:**

```go
result, err := filter.Unique([]int{1, 2, 2, 3})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [1 2 3]
```

### Join

Joins the elements of a slice into a single string with a specified separator.

**Example:**

```go
result, err := filter.Join([]string{"hello", "world"}, " ")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: "hello world"
```

### First

Returns the first element of a slice.

**Example:**

```go
result, err := filter.First([]string{"first", "second", "third"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: "first"
```

### Last

Returns the last element of a slice.

**Example:**

```go
result, err := filter.Last([]string{"one", "two", "last"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: "last"
```

### Random

Selects a random element from a slice.

**Example:**

```go
result, err := filter.Random([]int{1, 2, 3, 4, 5})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Random element: %v\n", result)
```

### Reverse

Reverses the order of elements in a slice.

**Example:**

```go
result, err := filter.Reverse([]int{1, 2, 3})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [3 2 1]
```

### Shuffle

Randomly rearranges the elements of the slice.

**Example:**

```go
original := []int{1, 2, 3, 4, 5}
shuffled, err := filter.Shuffle(original)
if err != nil {
    log.Fatal(err)
}
fmt.Println(shuffled) // Outputs a shuffled version of [1 2 3 4 5]
```

### Size

Returns the size (length) of a slice.

**Example:**

```go
size, err := filter.Size([]string{"one", "two", "three"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(size) // Outputs: 3
```

### Max

Finds and returns the maximum value from a slice of numbers.

**Example:**

```go
result, err := filter.Max([]float64{1.2, 3.4, 2.5})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 3.4
```

### Min

Finds and returns the minimum value from a slice of numbers.

**Example:**

```go
result, err := filter.Min([]float64{-1, 0, 2})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: -1
```

### Sum

Calculates the sum of all elements in a slice of numbers.

**Example:**

```go
result, err := filter.Sum([]float64{1, 2, 3})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 6
```

### Average

Calculates the average value of elements in a slice of numbers.

**Example:**

```go
result, err := filter.Average([]float64{1, 2, 3, 4})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 2.5
```

### Map

Extracts a slice of values for a specified key from each map in the input slice.

**Example:**

```go
input := []map[string]interface{}{
    {"key": "value1"},
    {"key": "value2"},
}
result, err := filter.Map(input, "key")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: ["value1" "value2"]
```
