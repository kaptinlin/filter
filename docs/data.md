# Data Functions in the `filter` Package

The `filter` package provides powerful functions for extracting nested values from complex data structures in Go, including maps, slices, arrays, structs, pointers, and interfaces.

## Functions

### Extract

Retrieves a nested value from any supported data structure using a dot-separated key path. This function supports complex nested structures and provides comprehensive error handling.

**Supported Data Types:**
- Maps (`map[string]interface{}` and similar)
- Slices and arrays (`[]interface{}`, `[N]Type`, multi-dimensional arrays)
- Structs (with JSON tags or exported field names)
- Pointers to any of the above
- Interfaces containing any of the above
- Complex nested combinations

**Error Handling:**
- `ErrKeyNotFound`: Key or field doesn't exist
- `ErrIndexOutOfRange`: Array/slice index out of bounds
- `ErrInvalidKeyType`: Invalid navigation (e.g., into primitive types)
- `ErrUnsupportedType`: Input is nil or unsupported

**Example with Map:**

```go
data := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "John Doe",
        "profile": map[string]interface{}{
            "address": map[string]interface{}{
                "city": "New York",
                "coordinates": []float64{40.7128, -74.0060},
            },
        },
    },
}

// Extract nested string
city, err := filter.Extract(data, "user.profile.address.city")
if err != nil {
    log.Fatal(err)
}
fmt.Println(city) // Outputs: "New York"

// Extract array element
lat, err := filter.Extract(data, "user.profile.address.coordinates.0")
if err != nil {
    log.Fatal(err)
}
fmt.Println(lat) // Outputs: 40.7128
```

**Example with Slice:**

```go
data := []interface{}{
    []interface{}{"zero", "one", "two"},
    []interface{}{0, 1, 2},
    map[string]interface{}{
        "nested": "value",
    },
}

// Access nested slice element
value, err := filter.Extract(data, "0.2") // Third element of first slice
if err != nil {
    log.Fatal(err)
}
fmt.Println(value) // Outputs: "two"

// Access nested map in slice
nested, err := filter.Extract(data, "2.nested")
if err != nil {
    log.Fatal(err)
}
fmt.Println(nested) // Outputs: "value"
```

**Example with Struct:**

```go
type Address struct {
    Street   string `json:"street"`
    City     string `json:"city"`
    PostCode int    `json:"post_code"`
}

type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Address Address  `json:"address"`
    Tags    []string `json:"tags"`
}

person := Person{
    Name: "Alice Smith",
    Age:  30,
    Address: Address{
        Street:   "123 Main St",
        City:     "Boston",
        PostCode: 12345,
    },
    Tags: []string{"developer", "golang", "backend"},
}

// Extract struct field
name, err := filter.Extract(person, "name")
if err != nil {
    log.Fatal(err)
}
fmt.Println(name) // Outputs: "Alice Smith"

// Extract nested struct field
city, err := filter.Extract(person, "address.city")
if err != nil {
    log.Fatal(err)
}
fmt.Println(city) // Outputs: "Boston"

// Extract array element from struct
tag, err := filter.Extract(person, "tags.1")
if err != nil {
    log.Fatal(err)
}
fmt.Println(tag) // Outputs: "golang"
```

**Example with Pointers:**

```go
type Department struct {
    Name     string `json:"name"`
    Location string `json:"location"`
}

type Employee struct {
    Name       string      `json:"name"`
    Department *Department `json:"department"`
}

employee := Employee{
    Name: "Bob Johnson",
    Department: &Department{
        Name:     "Engineering",
        Location: "Building A",
    },
}

// Extract through pointer
dept, err := filter.Extract(employee, "department.name")
if err != nil {
    log.Fatal(err)
}
fmt.Println(dept) // Outputs: "Engineering"
```

**Example with Multi-dimensional Arrays:**

```go
// 2D array example
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Extract element from 2D array
value, err := filter.Extract(matrix, "1.2") // Row 1, Column 2
if err != nil {
    log.Fatal(err)
}
fmt.Println(value) // Outputs: 6
```

**Example with Complex Nested Structure:**

```go
complexData := map[string]interface{}{
    "company": map[string]interface{}{
        "departments": []map[string]interface{}{
            {
                "name": "Engineering",
                "employees": []map[string]interface{}{
                    {"name": "Alice", "skills": []string{"Go", "Python"}},
                    {"name": "Bob", "skills": []string{"JavaScript", "React"}},
                },
            },
            {
                "name": "Marketing",
                "employees": []map[string]interface{}{
                    {"name": "Carol", "skills": []string{"SEO", "Content"}},
                },
            },
        },
    },
}

// Extract deeply nested value
skill, err := filter.Extract(complexData, "company.departments.0.employees.1.skills.0")
if err != nil {
    log.Fatal(err)
}
fmt.Println(skill) // Outputs: "JavaScript"
```

**Error Handling Examples:**

```go
data := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "John",
        "age":  30,
    },
}

// Key not found
_, err := filter.Extract(data, "user.email")
if errors.Is(err, filter.ErrKeyNotFound) {
    fmt.Println("Email field not found")
}

// Invalid navigation into primitive
_, err = filter.Extract(data, "user.age.invalid")
if errors.Is(err, filter.ErrInvalidKeyType) {
    fmt.Println("Cannot navigate into primitive value")
}

// Index out of range
slice := []string{"a", "b", "c"}
_, err = filter.Extract(slice, "5")
if errors.Is(err, filter.ErrIndexOutOfRange) {
    fmt.Println("Index out of range")
}
```
