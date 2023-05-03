package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("tmp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	b := []int{1, 2, 3, 4}
	fmt.Println("dcl:", b)

}
