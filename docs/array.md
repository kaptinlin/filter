# Array Functions in the `filter` Package

The `filter` package provides a collection of functions designed to manipulate and analyze slices in Go effectively.

## Functions

### Unique

Removes duplicate elements from a slice while preserving first-seen order. Works on comparable element types (numbers, strings, booleans, pointers, simple structs). Slices, maps, and other non-comparable element types return an error. Use `UniqueBy` to deduplicate records by a property.

**Example:**

```go
result, err := filter.Unique([]int{1, 2, 2, 3})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [1 2 3]
```

### UniqueBy

Removes duplicate elements by the value at a dot-separated key while preserving first-seen order. Missing or unreachable keys return an error. Extracted key values may be non-comparable; equality follows the package's dynamic value equality rules.

**Example:**

```go
products := []any{
    map[string]any{"handle": "shirt", "title": "Red Shirt"},
    map[string]any{"handle": "shoe", "title": "Blue Shoe"},
    map[string]any{"handle": "shirt", "title": "Green Shirt"},
}
result, err := filter.UniqueBy(products, "handle")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Keeps Red Shirt, Blue Shoe
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

Selects a random element from a slice. Each call picks independently, so repeated calls can return different results. Empty slices return an error. For deterministic output in tests, use `filter.RandomWithRand(filter.SeededRand(s1, s2), input)`.

**Example:**

```go
result, err := filter.Random([]int{1, 2, 3, 4, 5})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Random element: %v\n", result)

stable, err := filter.RandomWithRand(filter.SeededRand(1, 2), []int{1, 2, 3, 4, 5})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Stable random element: %v\n", stable)
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

Returns a new slice with elements rearranged in random order. Each invocation may produce a fresh order. For a stable order in tests, call `filter.ShuffleWithRand(filter.SeededRand(s1, s2), input)`.

**Example:**

```go
original := []int{1, 2, 3, 4, 5}
shuffled, err := filter.Shuffle(original)
if err != nil {
    log.Fatal(err)
}
fmt.Println(shuffled) // Outputs a shuffled version of [1 2 3 4 5]

stable, err := filter.ShuffleWithRand(filter.SeededRand(1, 2), original)
if err != nil {
    log.Fatal(err)
}
fmt.Println(stable) // Same order every run for the same seed
```

### Size

Returns the size (length) of a slice, array, map, or string. Strings use UTF-8 rune count, matching `Length`.

**Example:**

```go
size, err := filter.Size([]string{"one", "two", "three"})
if err != nil {
    log.Fatal(err)
}
fmt.Println(size) // Outputs: 3

size, _ = filter.Size("a界b")
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

### SumBy

Calculates the sum of numeric values extracted from each element at a dot-separated key. Missing keys and non-numeric values return an error.

**Example:**

```go
products := []any{
    map[string]any{"price": 50},
    map[string]any{"price": "30.5"},
    map[string]any{"price": 10.25},
}
result, err := filter.SumBy(products, "price")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: 90.75
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

Extracts a slice of values for a specified key from each element in the input slice. Output length always equals input length: when the key cannot be extracted from an element (missing key, missing index, type mismatch), the corresponding output is `nil` and no error is returned.

**Example:**

```go
input := []map[string]interface{}{
    {"key": "value1"},
    {"key": "value2"},
    {"other": "x"},
}
result, err := filter.Map(input, "key")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // Outputs: [value1 value2 <nil>]
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

Filters a slice, keeping elements where the given property equals the given value. If value is omitted, keeps elements where the property is truthy — only `nil` and `false` are falsy, so `0`, `""`, and empty collections are kept.

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

Returns the first element in a slice where the given property equals the given value. When nothing matches, returns an error matching `errors.Is(err, filter.ErrNotFound)`.

**Example:**

```go
products := []any{
    map[string]any{"handle": "shoes", "price": 50.0},
    map[string]any{"handle": "shirt", "price": 30.0},
}
result, _ := filter.Find(products, "handle", "shirt")
// Returns: map[handle:shirt price:30]

_, err := filter.Find(products, "handle", "hat")
if errors.Is(err, filter.ErrNotFound) {
    // no matching product
}
```

### FindIndex

Returns the 0-based index of the first element where the given property equals the given value. Returns `-1` (with no error) when nothing matches.

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
