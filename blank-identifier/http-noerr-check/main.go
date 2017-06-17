package main

import "net/http"
import "io/ioutil"
import "fmt"

func main() {
	// Call the http.Get but ignore errors with blank identifier.
	// Both functions return an error, but we don't care
	res, _ := http.Get("http://www.google.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", page)
}
