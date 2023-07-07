package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string, 3)
	fmt.Println("init:", m)

	m2 := map[string]int{
		"a": 5,
	}
	fmt.Println("dcl:", m2)

	m["key1"] = "value1"
	m["key1"] = "value1"
	m["key2"] = "value1"
	fmt.Println("set:", m)

	// zero_value for string is ""
	fmt.Println("valueNotExists:", m["k;;r"])

	// len of map
	fmt.Println("len of map", len(m))
	m["key1"] = ""
	fmt.Println("len of map", len(m))
	delete(m, "key1")
	delete(m, "key3")
	fmt.Println("len of map", len(m))

	val, exists := m["key1"]
	fmt.Printf("val: %v, exists: %v\n", val, exists)
	m["key1"] = ""
	val, exists = m["key1"]
	fmt.Printf("val: %v, exists: %v\n", val, exists)

	fmt.Println("check pass map is a reference")
	m1 := map[int]int{1: -1}
	fmt.Printf("m1 address: %p\n", m1)
	// m1 是指向对象的指针，指针值传递
	changeMapVal(m1)
	fmt.Printf("map after change:%v\n", m1)
}

func changeMapVal(m map[int]int) {
	fmt.Printf("changeMapVal m1 address: %p\n", m)
	m[1] = 1
}
