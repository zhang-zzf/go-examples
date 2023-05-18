package main

import (
	"fmt"
	"time"
)

func main() {

	// 定速
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.NewTicker(200 * time.Millisecond)
	for req := range requests {
		<-limiter.C
		fmt.Println(time.Now(), "->", req)
	}

	// 令牌桶
	burstyLimiter := make(chan time.Time, 5)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 200 millis 放入一个令牌
	go func() {
		for token := range time.Tick(time.Second) {
			burstyLimiter <- token
		}
	}()
	requests = make(chan int, 10)
	for i := 0; i < 10; i++ {
		requests <- i
	}
	close(requests)
	for req := range requests {
		token := <-burstyLimiter
		fmt.Println(token, "->", req)
	}

}
