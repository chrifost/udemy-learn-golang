package main

import "fmt"

func main() {
	//Simple example of hwo to turn a string into a slice of byte
	myString := "Chris was here"

	fmt.Println(myString)
	fmt.Println([]byte(myString))
	fmt.Printf("%c\n", []byte(myString))

	//NOTE - This is what a lot of go functions want to pass in/out
}
