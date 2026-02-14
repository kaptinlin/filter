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
	fmt.Println(filter.Truncate("Hello, World!", 5))
	fmt.Println(filter.Truncate("Hi", 5))
	// Output:
	// Hello...
	// Hi
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
