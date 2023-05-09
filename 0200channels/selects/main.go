package main

import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "a"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "b"
	}()

	for i := 0; i < 1; i++ {
		select {
		case msg1 := <-c1:
			println(msg1)
		case msg2 := <-c2:
			println(msg2)
		}
	}

	time.Sleep(time.Second * 5)
}
