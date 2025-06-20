package main

import "fmt"

const Pi = 3.14159              // constant of type float64
const Language = "GoLang"       // string constant
const Version = 1.21            // float constant

// Combining constants to define a new constant
// const FullName = Language + " v" + fmt.Sprint(Version) // ❌ This will cause a compilation error

// Instead, you must use only constant expressions:
const FullNameFixed = Language + " v1.21" // ✅ correct constant expression

func main2() {
	fmt.Println("Pi is:", Pi)
	fmt.Println("Language:", Language)
	fmt.Println("Version:", Version)
	fmt.Println("FullNameFixed:", FullNameFixed)

	// ❌ Uncommenting the next line will cause a compile-time error
	// Pi = 3.15

	// ✅ Constants are evaluated at compile-time and embedded into the binary.
	fmt.Println("🔍 Constants are compiled directly into the code, so Pi is hardcoded at compile time.")

	// You can use constants in expressions
	const Radius = 5
	const Area = Pi * Radius * Radius
	fmt.Println("Area of circle with radius 5 is:", Area)

	// You can group constants too
	const (
		Red   = "#FF0000"
		Green = "#00FF00"
		Blue  = "#0000FF"
	)
	fmt.Println("RGB Colors:", Red, Green, Blue)
}