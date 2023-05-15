package main

import "time"

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	println("timer1 fired")
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		println("timer2 fired")
	}()

	timer2.Stop()
	println("timer2 stopped")
	time.Sleep(time.Second * 2)
}
