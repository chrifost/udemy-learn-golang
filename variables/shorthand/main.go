package main

import "fmt"

func main() {

	// Shorthand notation; declares and initializes variable; type is inferred.
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "This is a sentence"
	f := 'c'
	g := `o`

	// Print the values
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)

	// Print the types for each
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

}
