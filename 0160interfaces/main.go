package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("g.area():", g.area())
	fmt.Println("g.perimeter()", g.perimeter())
}

func main() {
	r := rect{5, 2}
	measure(&r)
	c := circle{12.5}
	measure(&c)
	c2 := new(circle)
	c2.radius = 1.25
	measure(c2)
}
