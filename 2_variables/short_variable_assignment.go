package main

import (
	"fmt"
)

func main1() {
	// 🔹 SHORT VARIABLE DECLARATION
	// := is used to declare and initialize a variable in one step.
	// Go infers the type based on the right-hand side literal.
	name := "Dhanush"       // inferred as string
	age := 22               // inferred as int (default integer type)
	pi := 3.14              // inferred as float64 (default float type)
	isStudent := true       // inferred as bool
	initial := 'D'          // inferred as rune (int32 under the hood)

	fmt.Println("Short declaration values:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Pi:", pi)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Initial (rune):", initial)

	// 🔹 SAME LINE MULTIPLE DECLARATION
	// You can declare multiple variables in a single line.
	city, country := "Chennai", "India"
	fmt.Println("City:", city, "| Country:", country)

	// 🔹 TYPE CASTING / CONVERSION
	// Go is strictly typed — you must explicitly convert types.
	var a int = 42
	var b float64 = float64(a) // converting int to float64
	var c uint = uint(b)       // converting float64 to uint

	fmt.Println("Type casting examples:")
	fmt.Println("Int to Float64:", b)
	fmt.Println("Float64 to Uint:", c)

	// 🔹 DEFAULT TYPES (to stick to unless hyper-optimizing):
	// int     → default for integers
	// float64 → default for floating point
	// string  → default for text
	// bool    → true/false values
	// rune    → represents a Unicode character (alias for int32)
	// byte    → alias for uint8 (used for raw data like files, buffers)

	// Example using default types:
	var score int = 90           // prefer int unless memory-constrained
	var percentage float64 = 90.5 // prefer float64 for precision
	var grade string = "A"
	var passed bool = true

	fmt.Println("Default types:")
	fmt.Println("Score (int):", score)
	fmt.Println("Percentage (float64):", percentage)
	fmt.Println("Grade (string):", grade)
	fmt.Println("Passed (bool):", passed)

	// 🔹 HYPER OPTIMIZATION CASE (when needed)
	// Use int8, uint16, float32, etc., only when you need memory optimization,
	// for example in embedded systems, binary protocols, etc.
	var smallInt int8 = 100
	var tinyFloat float32 = 3.14

	fmt.Println("Optimized types:")
	fmt.Println("SmallInt (int8):", smallInt)
	fmt.Println("TinyFloat (float32):", tinyFloat)
	main2()
}