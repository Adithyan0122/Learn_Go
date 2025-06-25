package main

import "fmt"

// Define a struct
type Circle struct {
	radius float64
}

// Define a method on Circle to calculate area
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Define another method to calculate circumference
func (c Circle) circumference() float64 {
	return 2 * 3.14 * c.radius
}

func main4() {
	// Create a Circle instance
	c := Circle{radius: 5}

	// Call methods on the struct
	fmt.Println("Radius:", c.radius)
	fmt.Println("Area:", c.area())                   // Calling method
	fmt.Println("Circumference:", c.circumference()) // Calling another method
}
