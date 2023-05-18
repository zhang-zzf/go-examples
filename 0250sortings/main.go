package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"b", "z", "x", "a"}
	sort.Strings(strings)
	fmt.Println(strings)
	ss := sort.StringSlice(strings)
	ss.Sort()
	fmt.Println(ss.Len(), ss.Search("a"), ss.Search("y"))
	reverse := sort.Reverse(ss)
	sort.Sort(reverse)
	fmt.Println(reverse)

}
