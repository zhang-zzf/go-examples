package main

import "time"

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		c1 <- "1"
	}()

	select {
	case <-c1:
		println("c1 returned")
	case <-time.After(time.Second * 2):
		println("c1 timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "2"
	}()

	select {
	case <-c2:
		println("c2 returned.")
	case <-time.After(time.Second * 2):
		println("c2 timeout")
	}
}
