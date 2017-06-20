package main

import "fmt"

//simple const, global scope
const p string = "death and taxes"

//Declare a group of constants
const (
	z string  = "group of const"
	b         = 42
	d float32 = 3.14
)

func main() {
	//simple contant, local scope
	const q = 42
	a := 0
	a = 10

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("a - ", a)

	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", q)

}

// constants can be delclared with out wihout the type.  Obviously you cant assing
// new values to const, i.e. q = 5
