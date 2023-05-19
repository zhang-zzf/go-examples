package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var p = fmt.Println

func main() {
	bytes, err := os.ReadFile("/tmp/a.txt")
	checkErr(err)
	p(string(bytes))
	var file *os.File
	file, err = os.Open("/tmp/a.txt")
	checkErr(err)
	p(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err = scanner.Err(); err != nil {
		os.Exit(1)
	}
	// create file
	os.WriteFile("a.txt", []byte(""), 0644)

}
