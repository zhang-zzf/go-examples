package main

import "fmt"

/*
for is Go’s only looping construct. Here are some basic types of for loops.
*/
func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 5; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

}
