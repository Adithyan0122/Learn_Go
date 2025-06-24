package main

import "fmt"

//function signature
func concat(s1 string, s2 string) string{
	return s1 + s2
}

//if multiple arguments are of same type, the type only needs to be decalared after the last one
func add(x, y int) int{
	return x + y
}

func main(){
	fmt.Println(concat("Hello", " World!"))
	fmt.Println(add(1, 3))
	main2()
}