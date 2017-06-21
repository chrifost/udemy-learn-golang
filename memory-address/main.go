package main

import "fmt"

func main() {
	a := 43

	// Simply show value and memory address value is stored at
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)

	//Read a entered value and store it

	var meters float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&meters)
	fmt.Println("You entered - ", meters)
}
