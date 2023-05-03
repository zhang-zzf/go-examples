package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)
	var b, c = 1, 2
	fmt.Println(b, c)
	var d int
	// d will be initialized to 0
	fmt.Println(d)
	// bl will be false
	var bl bool
	fmt.Println(bl)
	// := syntax is shorthand for declaring and initializing a variable
	str := "apple"
	fmt.Println(str)
}
