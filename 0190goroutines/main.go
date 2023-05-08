package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, i)
	}
}

func main() {

	f("direct")

	go f("goroutine1")
	go f("goroutine2")

	go func(from string) {
		fmt.Println(from)
	}("going")

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
