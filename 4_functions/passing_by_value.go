/* 
varibales in go are passed by value (except for few data types).
Pass by Value means that when a variable is passed into a function, 
that function recieves a copy of the variable.
The function is unable to mutate the caller's data.
*/

package main 

import "fmt"

func increment(x int){
	x++
}

func main2(){
	x := 5 

	increment(x)

	fmt.Println(x)
	//still prints 5, 
	// because the increment function 
	// received a copy of x
	main3()
}