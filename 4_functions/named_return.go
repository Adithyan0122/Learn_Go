/*
Named Return and Naked Return in Go:

- Named return values are used by specifying variable names
  in the function's return type declaration.

  For example:
  func example(a int) (result int)

- Inside the function, we can directly assign values to
  these named return variables.

- A "naked return" means using the `return` statement
  without specifying any values.

- When a naked return is used, Go automatically returns
  the current values of the named return variables.

- This improves readability for small, simple functions.
  However, avoid naked returns in long functions as it
  can reduce code clarity.
*/

package main

import "fmt"

func calculateSumAndProduct(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b

	return
}

func main4() {
	a, b := 4, 5
	sum, product := calculateSumAndProduct(a, b)
	fmt.Printf("Sum:%d, Product:%d\n", sum, product)
}
