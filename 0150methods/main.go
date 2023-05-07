package main

import "fmt"

// Go supports methods defined on struct types.

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

// func (r *rect) perimeter() int {
func (r rect) perimeter() int {
	return 2 * (r.height + r.width)
}

func main() {

	r := &rect{5, 2}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

}
