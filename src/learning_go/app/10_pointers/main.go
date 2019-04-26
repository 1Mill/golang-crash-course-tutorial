package main

import "fmt"

func main() {
	a := 5
	b := &a // Point to memory address of a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // *[type] denotes pointer type

	// Use * to read val from address
	// & is the inverse to * through bijection. 
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
