package main

import "fmt"

func main() {
	ids := []int{33, 234, 345, 35, 345345}

	// Loop through ids using range
	for i,id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	// _ (underscore) denotes no using that param. In this case  (index, item) => (_, item)
	for _,id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _,element := range ids {
		sum += element
	}
	fmt.Println("Sum:", sum)

	// Range with map instead of slice
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com"}
	for key,value := range emails {
		fmt.Printf("%s's email is %s\n", key, value)
	}
	for key := range emails {
		fmt.Println("Name: " + key)
	}
}
