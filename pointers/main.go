package main

import "fmt"

// accept a memory address and update the value stored at that address
func zero(z *int) {
	fmt.Println("z's address is - ", z)
	fmt.Println("z has value - ", *z)

	*z = 0

	fmt.Println("z has new value - ", *z)

}

func main() {

	// print variable and it address
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's address is - ", &a)

	// Now assing a to a pointer
	// "b is a pointer to an int"
	// NOTE - This is actually:
	//   var b *int = &a
	// However, since we dont' know what a is, let go figure it out.  Go will
	// infer we need a *int based on this syntax
	var b = &a

	fmt.Println("b's address is - ", b)
	fmt.Println("b has value - ", *b)
	// C/C++ 101 - just dereference b with * to get the value of b, not the memory address

	// Now lets update the value stored in b's memory address
	// So just like C/C++ - update the value of stored at b, not where it points
	// Since b was pointing to a.. we actually updated a as well.
	*b = 42
	fmt.Println("a has new value - ", a)
	fmt.Println("b has new value - ", *b)

	//NOTE - Everyting is go is PASS BY VALUE.  When we pass a memory address,
	// we are passing a value

	// Example of how to use all this with function
	x := 5
	fmt.Println("x's address is - ", &x)
	fmt.Println("x has value - ", x)

	// call function with address of our memroy location
	zero(&x)
	fmt.Println("x has new value - ", x)

}
