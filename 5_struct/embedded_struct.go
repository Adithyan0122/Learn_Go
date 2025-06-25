package main

import "fmt"

// Define a base struct
type Address2 struct {
	city    string
	zipCode int
}

// Define a nested struct
type PersonNested struct {
	name     string
	age      int
	address2 Address2 // Nested: accessed via person.address.city
}

// Define an embedded struct
type PersonEmbedded struct {
	name     string
	age      int
	Address2 // Embedded: fields promoted directly into PersonEmbedded
}

func main3() {
	// Using Nested Struct
	nested := PersonNested{
		name: "Alice",
		age:  30,
		address2: Address2{
			city:    "New York",
			zipCode: 10001,
		},
	}

	// Using Embedded Struct
	embedded := PersonEmbedded{
		name: "Bob",
		age:  35,
		Address2: Address2{
			city:    "Los Angeles",
			zipCode: 90001,
		},
	}

	// Accessing nested fields
	fmt.Println("Nested Struct:")
	fmt.Println("Name:", nested.name)
	fmt.Println("City:", nested.address2.city) // Requires nested.address
	fmt.Println("Zip Code:", nested.address2.zipCode)

	// Accessing embedded fields
	fmt.Println("\nEmbedded Struct:")
	fmt.Println("Name:", embedded.name)
	fmt.Println("City:", embedded.city) // Direct access due to promotion
	fmt.Println("Zip Code:", embedded.zipCode)
	main4()
}
