package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan string)

	go func(msg chan string) {
		time.Sleep(time.Second * 1)
		msg <- "Hello, World"
		fmt.Println("message sent")
	}(message)

	start := time.Now()
	fmt.Println("start receiving msg:", start)
	fmt.Println(<-message)
	now := time.Now()
	fmt.Println("start receiving msg:", now)

	// fatal error: all goroutines are asleep - deadlock!
	// send will be blocked if there is no receiver
	fmt.Println("start send message 2")
	message <- "2"
	fmt.Println("sent message 2")

}
