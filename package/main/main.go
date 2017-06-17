package main

import (
	"fmt"

	"github.com/chrifost/udemy-learn-golang/package/stringutil"
)

func main() {
	//What is the value of our variable?
	fmt.Println(stringutil.MyName)

	//Print the lenght of the variable
	fmt.Println(stringutil.MyName, "is", stringutil.Reverse(stringutil.MyName), "characters")
}
