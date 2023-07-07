package main

import "fmt"

func main() {
	var s = make([]string, 3)
	fmt.Println("tmp:", s, len(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s, len(s))
	fmt.Printf("type: %T", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s, len(s))

	cpy := make([]string, len(s))
	copy(cpy, s)
	fmt.Println("copied: ", cpy)

	// slice
	l := s[4:]
	fmt.Println("sliced:", l)
	l[0] = "z"
	fmt.Println("updated slice:", l)
	// [a b c d z f]
	// slice is just a view of the origin
	fmt.Println("origin slice after update:", s)

	// t is also a Slices
	t := []int{1, 2, 3}
	t = append(t, 4)
	fmt.Println("dcl: ", t)

	// t2 is also a Slices
	// len(t2) is 0
	var t2 []int
	// runtime error
	// t2[0] = 1
	t2 = append(t2, 6)
	fmt.Println("dcl2:", t2, len(t2), cap(t2))

	var t3 [5]int
	fmt.Println("dcl3:", t3, len(t3))
	// error: t3 is an array
	// t3 = append(t3, 6)

	t4 := [...]int{1, 2}
	fmt.Println("array dcl:", t4)
	// error t4 is an array
	// t4 = append(t4, 3)

	s2 := make([]int, 4)
	s2[2] = 2
	fmt.Println("s2: ", s2, len(s2), cap(s2))
	s2 = append(s2, 6)
	fmt.Println("s2 after append: ", s2, len(s2), cap(s2))
	s2 = append(s2, 7)
	fmt.Println("s2 after append2: ", s2, len(s2), cap(s2))

	// 遍历
	squares := []int{1, 4, 9, 16, 25}
	aSlice1 := squares[:]
	for idx, val := range aSlice1 {
		if val == 9 {
			fmt.Println(idx)
		}
	}
	for i := 0; i < len(squares); i++ {
		if squares[i] == 0 {
			fmt.Println(i)
		}
	}

}
