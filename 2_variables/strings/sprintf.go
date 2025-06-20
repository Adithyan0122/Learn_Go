package main

import (
	"fmt"
)

func main2() {
	name := "Dhanush"
	age := 22
	score := 95.75

	// 🔹 fmt.Printf: Prints formatted output directly to the console
	fmt.Println("Using fmt.Printf:")
	fmt.Printf("Name: %s | Age: %d | Score: %.2f\n", name, age, score)
	// This will directly write to stdout (terminal)

	// 🔹 fmt.Sprintf: Returns a formatted string instead of printing it
	fmt.Println("\nUsing fmt.Sprintf:")
	message := fmt.Sprintf("Name: %s | Age: %d | Score: %.2f", name, age, score)
	// This returns a string which you can store, use, log, etc.

	// Print the message returned by Sprintf
	fmt.Println("Message returned by fmt.Sprintf:")
	fmt.Println(message)

	// 🔹 You can also use the Sprintf result in other contexts
	welcome := fmt.Sprintf("Welcome %s! Your age is %d.", name, age)
	fmt.Println("\nWelcome message built using fmt.Sprintf:")
	fmt.Println(welcome)
}