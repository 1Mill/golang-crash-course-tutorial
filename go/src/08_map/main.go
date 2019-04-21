package main

import "fmt"

func main() {
	// Def. map: (Key, Value) pair
	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete an entry
	delete(emails, "Bob")
	fmt.Println(emails)

	// Assign key, value on init
	someArray := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com"}
	fmt.Println(someArray)
}
