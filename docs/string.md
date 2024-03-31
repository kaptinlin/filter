# String Functions in the `filter` Package

The `filter` package provides a concise set of utilities designed to simplify common string manipulation tasks. 

## Functions

### Default

Returns a default value if the input string is empty.

**Example:**

```go
result := filter.Default("", "default value")
fmt.Println(result) // Outputs: "default value"
```

### Trim

Removes leading and trailing whitespace from the string.

**Example:**

```go
result := filter.Trim("  hello world  ")
fmt.Println(result) // Outputs: "hello world"
```

### Split

Divides a string into a slice of strings based on a specified delimiter.

**Example:**

```go
result := filter.Split("a,b,c", ",")
fmt.Println(result) // Outputs: ["a" "b" "c"]
```

### Replace

Substitutes all instances of a specified substring with another string.

**Example:**

```go
result := filter.Replace("hello world", "world", "go")
fmt.Println(result) // Outputs: "hello go"
```

### Remove

Eliminates all occurrences of a specified substring.

**Example:**

```go
result := filter.Remove("hello world", "world")
fmt.Println(result) // Outputs: "hello "
```

### Append

Adds characters to the end of a string.

**Example:**

```go
result := filter.Append("hello", " world")
fmt.Println(result) // Outputs: "hello world"
```

### Prepend

Adds characters to the beginning of a string.

**Example:**

```go
result := filter.Prepend("world", "hello ")
fmt.Println(result) // Outputs: "hello world"
```

### Length

Returns the number of characters in a string, accounting for UTF-8 encoding.

**Example:**

```go
result := filter.Length("hello")
fmt.Println(result) // Outputs: 5
```

### Upper

Converts all characters in a string to uppercase.

**Example:**

```go
result := filter.Upper("hello")
fmt.Println(result) // Outputs: "HELLO"
```

### Lower

Converts all characters in a string to lowercase.

**Example:**

```go
result := filter.Lower("HELLO")
fmt.Println(result) // Outputs: "hello"
```

### Titleize

Capitalizes the first letter of each word in a string.

**Example:**

```go
result := filter.Titleize("hello world")
fmt.Println(result) // Outputs: "Hello World"
```

### Capitalize

Capitalizes the first letter of a string.

**Example:**

```go
result := filter.Capitalize("hello")
fmt.Println(result) // Outputs: "Hello"
```

### Camelize

Converts a string to camelCase.

**Example:**

```go
result := filter.Camelize("hello_world")
fmt.Println(result) // Outputs: "helloWorld"
```

### Pascalize

Converts a string to PascalCase.

**Example:**

```go
result := filter.Pascalize("hello_world")
fmt.Println(result) // Outputs: "HelloWorld"
```

### Dasherize

Transforms a string into a lowercased, dash-separated format.

**Example:**

```go
result := filter.Dasherize("Hello World")
fmt.Println(result) // Outputs: "hello-world"
```

### Slugify

Converts a string into a URL-friendly "slug", ensuring it is safe for use in URLs and filenames by transliterating Unicode characters to ASCII, and replacing or removing special characters.

**Example:**

```go
text := filter.Slugify("Hellö Wörld хелло ворлд")
fmt.Println(text) // Outputs: "hello-world-khello-vorld"

someText := filter.Slugify("This & that")
fmt.Println(someText) // Outputs: "this-and-that"

anotherText := filter.Slugify("影師")
fmt.Println(anotherText) // Outputs: "ying-shi"
```

### Pluralize

Determines the singular or plural form of a word based on a numeric value.

**Example:**

```go
result := filter.Pluralize(2, "apple", "")
fmt.Println(result) // Outputs: "apples"
```

### Ordinalize

Converts a number to its ordinal English form.

**Example:**

```go
result := filter.Ordinalize(1)
fmt.Println(result) // Outputs: "1st"
```

### Truncate

Shortens a string to a specified length and appends "..." if it exceeds that length.

**Example:**

```go
result := filter.Truncate("hello world", 5)
fmt.Println(result) // Outputs: "hello..."
```

### TruncateWords

Truncates a string to a specified number of words, appending "..." if it exceeds that limit.

**Example:**

```go
result := filter.TruncateWords("hello beautiful world", 2)
fmt.Println(result) // Outputs: "hello beautiful..."
```
