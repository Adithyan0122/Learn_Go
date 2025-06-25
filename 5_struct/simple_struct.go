package main

import "fmt"

// Define a struct named person
type Person struct {
	name string
	age  int
}

func main() {
	//create a new person instance
	p1 := Person{
		name: "Alice",
		age:  30,
	}

	//Access struct fields
	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)

	//Modify struct fields
	p1.age = 32
	fmt.Println("Updated Age:", p1.age)
	main2()
}
