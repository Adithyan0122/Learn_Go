/*
Go does not allow unused variables—
they cause a compile-time error.

If a function returns a value you don't need,
you can ignore it using an underscore (_).

This is useful when you want to discard
unnecessary return values.
*/

package main

import "fmt"

func main3() {
	firstName, _ := getNames()
	fmt.Println("Welcome, ", firstName)
	main4()

}

func getNames() (string, string) {
	return "John", "Doe"
}
