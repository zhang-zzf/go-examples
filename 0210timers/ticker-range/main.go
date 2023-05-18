package main

import (
	"fmt"
	"time"
)

func main() {

	for ticker := range time.Tick(300 * time.Millisecond) {
		fmt.Println(time.Now(), "->", ticker)
		time.Sleep(time.Second)
	}

}
