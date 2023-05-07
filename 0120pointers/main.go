package main

import "fmt"

func main() {
	i, j := 1, 1
	zero_val(i)
	fmt.Println("i:", i)
	zero_ptr(&j)
	fmt.Println("j:", j)
}

func zero_ptr(ival *int) {
	*ival = 0
}

func zero_val(ival int) {
	ival = 0
}
