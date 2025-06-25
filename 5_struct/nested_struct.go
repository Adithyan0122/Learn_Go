package main

import "fmt"

// Address is a nested struct
type Address struct {
	city    string
	zipCode int
}

// Person struct contains Address as a nested struct
type Person2 struct {
	name    string
	age     int
	address Address
}

func main2() {
	// Empty initialization of a Person struct
	var p Person2

	// Assign values to fields after initialization
	p.name = "John"
	p.age = 25
	p.address.city = "New York"
	p.address.zipCode = 10001

	// Print the struct fields
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("Zip Code:", p.address.zipCode)
	main3()
}
