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

### Index

Returns the element at a specified index in a slice.

**Example:**

```go
result, err := filter.Index([]any{1, 2, 3}, 1)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 2
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

Returns the size (length) of a slice, array, or map.

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

### Sort

Sorts a slice in ascending order. If a key is provided, sorts a slice of maps/structs by that property. Numbers are compared numerically, strings lexicographically.

**Example:**

```go
result, err := filter.Sort([]any{"banana", "apple", "cherry"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [apple banana cherry]

// Sort by key
products := []any{
    map[string]any{"name": "Shoes", "price": 50.0},
    map[string]any{"name": "Shirt", "price": 30.0},
}
result, err = filter.Sort(products, "price")
fmt.Println(result) // Sorted by price ascending
```

### SortNatural

Sorts a slice case-insensitively. If a key is provided, sorts by that property case-insensitively.

**Example:**

```go
result, err := filter.SortNatural([]any{"Banana", "apple", "Cherry"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [apple Banana Cherry]
```

### Compact

Removes nil elements from a slice. If a key is provided, removes elements where the property is nil.

**Example:**

```go
result, err := filter.Compact([]any{"a", nil, "b", nil, "c"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [a b c]
```

### Concat

Combines two slices into one.

**Example:**

```go
result, err := filter.Concat([]any{"a", "b"}, []any{"c", "d"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [a b c d]
```

### Where

Filters a slice, keeping elements where the given property equals the given value. If value is omitted, keeps elements where the property is truthy (not nil and not false).

**Example:**

```go
products := []any{
    map[string]any{"name": "Shoes", "available": true},
    map[string]any{"name": "Shirt", "available": false},
    map[string]any{"name": "Pants", "available": true},
}

// Filter by value
result, _ := filter.Where(products, "available", true)
// Returns: Shoes, Pants

// Filter by truthy
result, _ = filter.Where(products, "available")
// Returns: Shoes, Pants (true is truthy, false is not)
```

### Reject

Filters a slice, removing elements where the given property equals the given value. If value is omitted, removes elements where the property is truthy. Inverse of Where.

**Example:**

```go
products := []any{
    map[string]any{"name": "Shoes", "available": true},
    map[string]any{"name": "Shirt", "available": false},
    map[string]any{"name": "Pants", "available": true},
}

result, _ := filter.Reject(products, "available", false)
// Returns: Shoes, Pants
```

### Find

Returns the first element in a slice where the given property equals the given value. Returns nil if not found.

**Example:**

```go
products := []any{
    map[string]any{"handle": "shoes", "price": 50.0},
    map[string]any{"handle": "shirt", "price": 30.0},
}
result, _ := filter.Find(products, "handle", "shirt")
// Returns: map[handle:shirt price:30]
```

### FindIndex

Returns the 0-based index of the first element where the given property equals the given value. Returns -1 if not found.

**Example:**

```go
products := []any{
    map[string]any{"handle": "shoes"},
    map[string]any{"handle": "shirt"},
    map[string]any{"handle": "pants"},
}
idx, _ := filter.FindIndex(products, "handle", "shirt")
fmt.Println(idx) // Outputs: 1
```

### Has

Returns true if any element in the slice has a property matching the given criteria. If value is provided, checks property equals value. If value is omitted, checks property is truthy.

**Example:**

```go
products := []any{
    map[string]any{"name": "Shoes", "available": true},
    map[string]any{"name": "Shirt", "available": false},
}

result, _ := filter.Has(products, "name", "Shoes")
fmt.Println(result) // Outputs: true

result, _ = filter.Has(products, "available")
fmt.Println(result) // Outputs: true
```
