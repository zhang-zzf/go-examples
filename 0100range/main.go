package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4}
	for i, v := range num {
		fmt.Printf("%v->%v\n", i, v)
	}

	m := map[int]int{1: 1, 2: 2}
	for k, v := range m {
		fmt.Printf("%v->%v\n", k, v)
	}

	str := "go"
	for i, c := range str {
		fmt.Printf("%v->%v\n", i, c)
	}
}
