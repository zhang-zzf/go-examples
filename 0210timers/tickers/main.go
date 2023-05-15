package main

import (
	"fmt"
	"time"
)

func main() {

	ticker1 := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for true {
			select {
			case <-done:
				return
			case t := <-ticker1.C:
				fmt.Println("receive ticker:", t)
			}
		}
	}()

	time.Sleep(time.Second * 4)
	// stop the ticker
	ticker1.Stop()
	fmt.Println("ticker1 stop")
	time.Sleep(time.Second * 8)
	done <- true
}
