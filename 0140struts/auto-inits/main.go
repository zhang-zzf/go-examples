package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// person will be zero initialized
	var person Person
	fmt.Println("person.init:", person)
}
