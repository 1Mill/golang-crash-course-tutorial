package main

import (
	"fmt"
	"math"
)

// Define our interface.
// * This outlines the methods that must be implemented by a struct 
// * (e.g. area, perimeter, etc.) to be clasified as that interfaces type.
// * For example, if Circle did not implement an area() method, an error
// * would be produced when you try to pass a Circle into getArea(s Shape)/
type Shape interface {
	area() float64
}

// Define structs
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// Define area functions
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Define common function that accepts shpae to get area
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
