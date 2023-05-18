package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unix := now.Unix()
	milli := now.UnixMilli()
	fmt.Println(unix, milli)
	anotherTime := time.Unix(unix, 0)
	anotherTime2 := time.Unix(0, now.UnixNano())
	// false
	fmt.Println(now.Equal(anotherTime))
	// true
	fmt.Println(now.Equal(anotherTime2))
}
