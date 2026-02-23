package filter_test

import (
	"fmt"

	"github.com/kaptinlin/filter"
)

func ExampleTrim() {
	result := filter.Trim("  hello  ")
	fmt.Println(result)
	// Output: hello
}

func ExampleCamelize() {
	fmt.Println(filter.Camelize("hello_world"))
	fmt.Println(filter.Camelize("foo-bar-baz"))
	// Output:
	// helloWorld
	// fooBarBaz
}

func ExamplePascalize() {
	fmt.Println(filter.Pascalize("hello_world"))
	fmt.Println(filter.Pascalize("foo-bar"))
	// Output:
	// HelloWorld
	// FooBar
}

func ExampleDasherize() {
	fmt.Println(filter.Dasherize("helloWorld"))
	fmt.Println(filter.Dasherize("FooBar"))
	// Output:
	// hello-world
	// foo-bar
}

func ExampleSlugify() {
	result := filter.Slugify("Hello World!")
	fmt.Println(result)
	// Output: hello-world
}

func ExampleTruncate() {
	fmt.Println(filter.Truncate("Hello, World!", 8))
	fmt.Println(filter.Truncate("Hi", 5))
	fmt.Println(filter.Truncate("Hello, World!", 10, "--"))
	// Output:
	// Hello...
	// Hi
	// Hello, W--
}

func ExamplePluralize() {
	fmt.Println(filter.Pluralize(1, "item", "items"))
	fmt.Println(filter.Pluralize(5, "item", "items"))
	// Output:
	// item
	// items
}

func ExampleOrdinalize() {
	fmt.Println(filter.Ordinalize(1))
	fmt.Println(filter.Ordinalize(2))
	fmt.Println(filter.Ordinalize(3))
	fmt.Println(filter.Ordinalize(11))
	// Output:
	// 1st
	// 2nd
	// 3rd
	// 11th
}

func ExampleExtract() {
	data := map[string]any{
		"user": map[string]any{
			"name": "Alice",
			"age":  30,
		},
	}
	name, _ := filter.Extract(data, "user.name")
	fmt.Println(name)
	// Output: Alice
}

func ExampleUnique() {
	result, _ := filter.Unique([]any{1, 2, 2, 3, 3, 3})
	fmt.Println(result)
	// Output: [1 2 3]
}

func ExampleJoin() {
	result, _ := filter.Join([]string{"a", "b", "c"}, ", ")
	fmt.Println(result)
	// Output: a, b, c
}

func ExampleAbs() {
	result, _ := filter.Abs(-5)
	fmt.Println(result)
	// Output: 5
}

func ExampleRound() {
	result, _ := filter.Round(3.14159, 2)
	fmt.Println(result)
	// Output: 3.14
}

func ExampleEscape() {
	fmt.Println(filter.Escape("<p>Hello & World</p>"))
	// Output: &lt;p&gt;Hello &amp; World&lt;/p&gt;
}

func ExampleEscapeOnce() {
	fmt.Println(filter.EscapeOnce("&lt;p&gt;already escaped&lt;/p&gt;"))
	fmt.Println(filter.EscapeOnce("1 < 2 & 3"))
	// Output:
	// &lt;p&gt;already escaped&lt;/p&gt;
	// 1 &lt; 2 &amp; 3
}

func ExampleStripHTML() {
	fmt.Println(filter.StripHTML("<p>Hello <b>World</b></p>"))
	// Output: Hello World
}

func ExampleTrimLeft() {
	fmt.Println(filter.TrimLeft("  hello  "))
	// Output: hello
}

func ExampleTrimRight() {
	fmt.Println(filter.TrimRight("  hello  "))
	// Output:   hello
}

func ExampleReplaceFirst() {
	fmt.Println(filter.ReplaceFirst("hello hello hello", "hello", "hi"))
	// Output: hi hello hello
}

func ExampleReplaceLast() {
	fmt.Println(filter.ReplaceLast("hello hello hello", "hello", "hi"))
	// Output: hello hello hi
}

func ExampleUrlEncode() {
	fmt.Println(filter.UrlEncode("hello world"))
	// Output: hello+world
}

func ExampleBase64Encode() {
	fmt.Println(filter.Base64Encode("hello world"))
	// Output: aGVsbG8gd29ybGQ=
}

func ExampleBase64Decode() {
	result, _ := filter.Base64Decode("aGVsbG8gd29ybGQ=")
	fmt.Println(result)
	// Output: hello world
}

func ExampleSort() {
	result, _ := filter.Sort([]any{"banana", "apple", "cherry"})
	fmt.Println(result)
	// Output: [apple banana cherry]
}

func ExampleCompact() {
	result, _ := filter.Compact([]any{"a", nil, "b", nil, "c"})
	fmt.Println(result)
	// Output: [a b c]
}

func ExampleConcat() {
	result, _ := filter.Concat([]any{"a", "b"}, []any{"c", "d"})
	fmt.Println(result)
	// Output: [a b c d]
}

func ExampleWhere() {
	products := []any{
		map[string]any{"name": "Shoes", "available": true},
		map[string]any{"name": "Shirt", "available": false},
		map[string]any{"name": "Pants", "available": true},
	}
	result, _ := filter.Where(products, "available", true)
	for _, p := range result {
		m := p.(map[string]any)
		fmt.Println(m["name"])
	}
	// Output:
	// Shoes
	// Pants
}

func ExampleDefault() {
	fmt.Println(filter.Default("", "fallback"))
	fmt.Println(filter.Default("value", "fallback"))
	fmt.Println(filter.Default(nil, "fallback"))
	fmt.Println(filter.Default(false, "fallback"))
	// Output:
	// fallback
	// value
	// fallback
	// fallback
}

func ExampleCapitalize() {
	fmt.Println(filter.Capitalize("hELLO"))
	fmt.Println(filter.Capitalize("hello world"))
	// Output:
	// Hello
	// Hello world
}
