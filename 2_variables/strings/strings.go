package main

import "fmt"

func main() {
	// Declare variables of different types
	name := "Dhanush"
	age := 22
	pi := 3.1415926535

	// 🔹 1. Default representation using %v (generic placeholder)
	fmt.Printf("Default formatting with %%v:\n")
	fmt.Printf("Name: %v, Age: %v, Pi: %v\n\n", name, age, pi)

	// 🔹 2. Interpolating a string using %s
	fmt.Printf("String interpolation with %%s:\n")
	fmt.Printf("Hello, %s!\n\n", name)

	// 🔹 3. Interpolating an integer using %d (decimal format)
	fmt.Printf("Integer interpolation with %%d:\n")
	fmt.Printf("%s is %d years old.\n\n", name, age)

	// 🔹 4. Interpolating a float with default precision using %f
	fmt.Printf("Float interpolation with %%f:\n")
	fmt.Printf("Default Pi value: %f\n\n", pi)

	// 🔹 5. Interpolating a float with specified precision: %.2f (2 decimal places)
	fmt.Printf("Float interpolation with precision using %%.2f:\n")
	fmt.Printf("Pi to 2 decimal places: %.2f\n", pi)
	fmt.Printf("Pi to 4 decimal places: %.4f\n\n", pi)

	// 🔹 6. Using fmt.Sprintf to create and return a formatted string
	message := fmt.Sprintf("%s is %.1f times cooler than average!", name, 1.618)
	fmt.Println("Using fmt.Sprintf to build a formatted string:")
	fmt.Println(message)

	// 🔹 7. Formatting boolean, character, and hexadecimal values
	isAwesome := true
	initial := 'D'
	number := 255

	fmt.Printf("\nOther formats:\n")
	fmt.Printf("Boolean: %t\n", isAwesome)     // true or false
	fmt.Printf("Character: %c\n", initial)     // Unicode character
	fmt.Printf("Hexadecimal: %x\n", number)    // hexadecimal (lowercase)
	main2()
}