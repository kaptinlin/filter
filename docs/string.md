# String Functions in the `filter` Package

The `filter` package provides a concise set of utilities designed to simplify common string manipulation tasks.

## Functions

### Default

Returns `defaultValue` if the input is `nil`, `false`, or an empty string. Accepts any type.

**Example:**

```go
result := filter.Default("", "fallback")
fmt.Println(result) // Outputs: "fallback"

result = filter.Default(nil, "fallback")
fmt.Println(result) // Outputs: "fallback"

result = filter.Default(false, "fallback")
fmt.Println(result) // Outputs: "fallback"

result = filter.Default("value", "fallback")
fmt.Println(result) // Outputs: "value"
```

### Trim

Removes leading and trailing whitespace from the string.

**Example:**

```go
result := filter.Trim("  hello world  ")
fmt.Println(result) // Outputs: "hello world"
```

### TrimLeft

Removes leading whitespace from a string. Equivalent to Liquid's `lstrip`.

**Example:**

```go
result := filter.TrimLeft("  hello  ")
fmt.Println(result) // Outputs: "hello  "
```

### TrimRight

Removes trailing whitespace from a string. Equivalent to Liquid's `rstrip`.

**Example:**

```go
result := filter.TrimRight("  hello  ")
fmt.Println(result) // Outputs: "  hello"
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

### ReplaceFirst

Replaces the first occurrence of a substring with another string.

**Example:**

```go
result := filter.ReplaceFirst("hello hello hello", "hello", "hi")
fmt.Println(result) // Outputs: "hi hello hello"
```

### ReplaceLast

Replaces the last occurrence of a substring with another string.

**Example:**

```go
result := filter.ReplaceLast("hello hello hello", "hello", "hi")
fmt.Println(result) // Outputs: "hello hello hi"
```

### Remove

Eliminates all occurrences of a specified substring.

**Example:**

```go
result := filter.Remove("hello world", "world")
fmt.Println(result) // Outputs: "hello "
```

### RemoveFirst

Removes the first occurrence of a substring.

**Example:**

```go
result := filter.RemoveFirst("hello hello hello", "hello ")
fmt.Println(result) // Outputs: "hello hello"
```

### RemoveLast

Removes the last occurrence of a substring.

**Example:**

```go
result := filter.RemoveLast("hello hello hello", " hello")
fmt.Println(result) // Outputs: "hello hello"
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

Capitalizes the first letter and lowercases the rest. Follows the Liquid `capitalize` behavior.

**Example:**

```go
result := filter.Capitalize("hELLO")
fmt.Println(result) // Outputs: "Hello"

result = filter.Capitalize("hello world")
fmt.Println(result) // Outputs: "Hello world"
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

result := filter.Pluralize(1, "%d message", "%d messages")
fmt.Println(result) // Outputs: "1 message"
```

### Ordinalize

Converts a number to its ordinal English form.

**Example:**

```go
result := filter.Ordinalize(1)
fmt.Println(result) // Outputs: "1st"
```

### Truncate

Shortens a string to a specified length (including the ellipsis). An optional ellipsis string can be provided (default `"..."`).

**Example:**

```go
result := filter.Truncate("Hello, World!", 8)
fmt.Println(result) // Outputs: "Hello..."

result = filter.Truncate("Hello, World!", 10, "--")
fmt.Println(result) // Outputs: "Hello, W--"

result = filter.Truncate("Hi", 5)
fmt.Println(result) // Outputs: "Hi"
```

### TruncateWords

Truncates a string to a specified number of words. An optional ellipsis string can be provided (default `"..."`).

**Example:**

```go
result := filter.TruncateWords("hello beautiful world", 2)
fmt.Println(result) // Outputs: "hello beautiful..."

result = filter.TruncateWords("hello beautiful world", 2, "--")
fmt.Println(result) // Outputs: "hello beautiful--"
```

### Escape

HTML-escapes a string, converting `<`, `>`, `&`, `"`, `'` to HTML entities.

**Example:**

```go
result := filter.Escape("<p>Hello & World</p>")
fmt.Println(result) // Outputs: "&lt;p&gt;Hello &amp; World&lt;/p&gt;"
```

### EscapeOnce

HTML-escapes a string without double-escaping existing entities. Already escaped entities like `&amp;` and `&lt;` are preserved.

**Example:**

```go
result := filter.EscapeOnce("&lt;p&gt;already escaped&lt;/p&gt;")
fmt.Println(result) // Outputs: "&lt;p&gt;already escaped&lt;/p&gt;"

result = filter.EscapeOnce("1 < 2 & 3")
fmt.Println(result) // Outputs: "1 &lt; 2 &amp; 3"
```

### StripHTML

Removes all HTML tags, script blocks, style blocks, and comments from the input.

**Example:**

```go
result := filter.StripHTML("<p>Hello <b>World</b></p>")
fmt.Println(result) // Outputs: "Hello World"

result = filter.StripHTML("before<script>alert('x')</script>after")
fmt.Println(result) // Outputs: "beforeafter"
```

### StripNewlines

Removes all newline characters (`\n`, `\r\n`, `\r`) from the input.

**Example:**

```go
result := filter.StripNewlines("hello\nworld")
fmt.Println(result) // Outputs: "helloworld"
```

### Slice

Extracts a substring or sub-slice. Negative offset counts from end. If length is omitted, returns a single character/element.

**Example:**

```go
// String slicing
result, _ := filter.Slice("hello", 1, 3)
fmt.Println(result) // Outputs: "ell"

result, _ = filter.Slice("hello", -3, 2)
fmt.Println(result) // Outputs: "ll"

// Array slicing
result, _ = filter.Slice([]any{1, 2, 3, 4}, 1, 2)
fmt.Println(result) // Outputs: [2 3]
```

### UrlEncode

Percent-encodes a string for use in URLs.

**Example:**

```go
result := filter.UrlEncode("hello world")
fmt.Println(result) // Outputs: "hello+world"

result = filter.UrlEncode("foo@bar.com")
fmt.Println(result) // Outputs: "foo%40bar.com"
```

### UrlDecode

Decodes a percent-encoded string.

**Example:**

```go
result, _ := filter.UrlDecode("hello+world")
fmt.Println(result) // Outputs: "hello world"
```

### Base64Encode

Encodes a string to standard Base64.

**Example:**

```go
result := filter.Base64Encode("hello world")
fmt.Println(result) // Outputs: "aGVsbG8gd29ybGQ="
```

### Base64Decode

Decodes a standard Base64 string.

**Example:**

```go
result, _ := filter.Base64Decode("aGVsbG8gd29ybGQ=")
fmt.Println(result) // Outputs: "hello world"
```
