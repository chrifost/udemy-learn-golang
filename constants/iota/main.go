package main

import "fmt"

// Declare some simple iota's.  First occurent always starts at 0, then increment
const (
	a = iota //0
	b = iota //1
	c = iota //2
)

// If you start a new const block, the iota resets to 0
const (
	d = iota //0
	e = iota //1
	f = iota //2
)

// Now do some bitwise operations; notice you can use blank identifier for assignment
const (
	_  = iota //0
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("KB - %d, %b \n", kb, kb)
	fmt.Printf("MB - %d, %b \n", mb, mb)
	fmt.Printf("MB = %d \n", (1024 * 1024))

}
