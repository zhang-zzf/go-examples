package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "1"
	messages <- "2"
	// fatal error: all goroutines are asleep - deadlock!
	// messages <- "3"

	println(<-messages, <-messages)
	// 	println(<-messages, <-messages, <-messages)
	// println(<-messages, <-messages, <-messages)
	fmt.Println("done")
}
