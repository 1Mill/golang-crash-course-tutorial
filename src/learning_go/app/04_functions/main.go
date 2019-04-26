package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name + "!"
}

func getSum(num1,num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("person"))
	fmt.Println(greeting("some name"))
	fmt.Println(greeting("another name"))

	fmt.Println(getSum(3,4))
}
