package main

import "fmt"

func main() {

	// Declare and init to default values

	var a int
	var b string
	var c float64
	var d bool

	// Print the values
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	// Print the types for each
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

}
