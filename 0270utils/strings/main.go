package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:", s.Contains("test", "st"))
	p("Contains:", s.Contains("test", "tt"))
	p("Count:", s.Count("test", "t"))
	p("Replace:", s.Replace("test", "t", "s", 2))

}
