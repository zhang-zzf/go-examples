package main

import "time"

func worker(done chan bool) {
	println("start...")
	time.Sleep(time.Second)
	done <- true
	println("done.")
}

func main() {
	done := make(chan bool)
	go worker(done)
	<-done
}
