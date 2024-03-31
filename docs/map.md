# Map Functions in the `filter` Package

The `filter` package offers a function for extracting nested values from maps, slices, and arrays in Go, simplifying access to nested data.

## Function

### Extract

Retrieves a nested value from a map, slice, or array using a dot-separated key path. 

**Example with Map:**

```go
data := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "John Doe",
        "address": map[string]interface{}{
            "city": "New York",
        },
    },
}
value, err := filter.Extract(data, "user.address.city")
if err != nil {
    log.Fatal(err)
}
fmt.Println(value) // Outputs: "New York"
```

**Example with Slice:**

```go
data := []interface{}{
    []interface{}{"zero", "one", "two"},
    []interface{}{0, 1, 2},
}
value, err := filter.Extract(data, "0.2") // Accessing the third element of the first slice
if err != nil {
    log.Fatal(err)
}
fmt.Println(value) // Outputs: "two"
```
