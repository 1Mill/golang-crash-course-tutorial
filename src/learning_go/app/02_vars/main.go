package main

import "fmt"

func main() {
	// Can use const or vars similar to javascript
	const name = "Some random name"
	const age = 1234

	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

	// Use short hand syntax
	new_name := "Another name"
	new_age := 12.34
	new_bool := false

	fmt.Println(new_name, new_age)
	fmt.Printf("%T\n", new_name)
	fmt.Printf("%T\n", new_age)
	fmt.Printf("%T\n", new_bool)

	// Can do multiple assignments
	a, b, c := "A new name that is long", 654.321, true
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
