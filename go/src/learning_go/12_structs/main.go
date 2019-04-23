package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
// Long form
// type Person struct {
// 	firstName string
// 	lastName string
// 	city string
// 	gender string
// 	age int
// }
// Short form
type Person struct {
	firstName, lastName, city, gender string
	age int
}

// Greeting metehod (value reciever: just accepts some value)
// Adds "virtual attribute" method to Person struct (e.g. instance method)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method; Pointer reciever: changes something
// Returns nothing, so no return value specificed
func (p *Person) hasBirthday() {
	p.age++
}

// getsMarried (pointer reciever)
func (p *Person) getsMarried(newLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = newLastName
	}
}


func main() {
	// Init person using struct
	person1 := Person{firstName: "Bob", lastName: "Smith", city: "Seattle", gender: "M", age: 12}
	person2 := Person{"Bob2", "Smith2", "Seattle", "M2", 123}
	person3 := Person{"Sam", "Something", "Yakima", "F", 1234}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	
	person1.firstName = "Something else"
	fmt.Println(person1.firstName)

	// Call value reciver function
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	// People getting married.
	person1.getsMarried(person3.lastName)
	fmt.Println(person1.greet())
	person3.getsMarried(person1.lastName)
	fmt.Println(person3.greet())

}
