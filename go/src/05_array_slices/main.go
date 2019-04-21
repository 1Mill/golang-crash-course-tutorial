package main

import "fmt"

func main() {
	// Arrays
	var fruitArray [2]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)

	// Declare and assign at the same time
	anotherArray := [2]string {"New Value", "Another new value"}
	fmt.Println(anotherArray)
	fmt.Println(anotherArray[1])

	// Slices
	fruitSlices := []string{"Apple", "Oranges", "Grapes"}
	fmt.Println(fruitSlices)
	fmt.Println(len(fruitSlices))
	fmt.Println(fruitSlices[1:2])
	fmt.Println(fruitSlices[1:3])
}
