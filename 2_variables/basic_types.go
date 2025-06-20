package main

import (
	"fmt"
)

func main() {
	// Boolean
	var isGoFun bool = true

	// String
	var greeting string = "Hello, Go!"

	// Signed Integers
	var a int = -42
	var b int8 = -8
	var c int16 = -1600
	var d int32 = -32000
	var e int64 = -64000000000

	// Unsigned Integers
	var ua uint = 42
	var ub uint8 = 8 // alias: byte
	var uc uint16 = 1600
	var ud uint32 = 32000
	var ue uint64 = 64000000000

	// Byte and Rune
	var myByte byte = 'A'       // byte is alias for uint8
	var myRune rune = '❤'       // rune is alias for int32

	// Floating Point Numbers
	var pi float32 = 3.14
	var gravity float64 = 9.80665

	// Complex Numbers
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 2 + 3i

	// Printing all variables
	fmt.Println("Boolean:", isGoFun)
	fmt.Println("String:", greeting)

	fmt.Println("\nSigned Integers:")
	fmt.Println("int:", a, "int8:", b, "int16:", c, "int32:", d, "int64:", e)

	fmt.Println("\nUnsigned Integers:")
	fmt.Println("uint:", ua, "uint8 (byte):", ub, "uint16:", uc, "uint32:", ud, "uint64:", ue)

	fmt.Println("\nByte and Rune:")
	fmt.Printf("byte: %v (char: %c)\n", myByte, myByte)
	fmt.Printf("rune: %v (char: %c)\n", myRune, myRune)

	fmt.Println("\nFloating Points:")
	fmt.Println("float32:", pi, "float64:", gravity)

	fmt.Println("\nComplex Numbers:")
	fmt.Println("complex64:", c1, "complex128:", c2)
	main1()
}