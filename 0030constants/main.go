package main

import "fmt"

const s = "constant"

const (
	c1 = "a"
	c2 = true
	c3 = 'a'
)

func main() {

	fmt.Println(s)

	const number = 1 << 8
	fmt.Println(number)

}
