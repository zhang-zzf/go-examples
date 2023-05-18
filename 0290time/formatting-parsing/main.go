package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	dateTime := now.Format(time.DateTime)
	fmt.Println(dateTime)
	rfc3339 := now.Format(time.RFC3339)
	fmt.Println(rfc3339)
	parse, _ := time.Parse(time.RFC3339, rfc3339)
	// false
	fmt.Println(now.Equal(parse))

}
