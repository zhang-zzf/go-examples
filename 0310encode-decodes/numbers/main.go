package main

import (
	"fmt"
	"strconv"
)

var p = fmt.Println

func main() {

	if float, err := strconv.ParseFloat("1.23", 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("parsed float:", float)
	}
	if integer, err := strconv.ParseInt("2.3", 0, 32); err != nil {
		// strconv.ParseInt: parsing "2.3": invalid syntax
		fmt.Println(err)
	} else {
		fmt.Println("parsed integer:", integer)
	}
	// -123
	i, err := strconv.ParseInt("-123", 0, 0)
	fmt.Printf("%T, %v, %v\n", i, i, err)
	i2, err2 := strconv.ParseInt("123", 0, 16)
	fmt.Printf("%T, %v, %v\n", i2, i2, err2)
	// 123
	p(strconv.ParseInt("+123", 0, 0))
	p(strconv.ParseInt("+123", 8, 0))
	p(strconv.ParseInt("+123", 16, 0))
	p(strconv.ParseInt("+123", 32, 0))
	// 2147483647 strconv.ParseInt: parsing "+123000000000": value out of range
	p(strconv.ParseInt("+123000000000", 32, 0))
	// 0x123 == 291
	p(strconv.ParseInt("0x123", 0, 0))
	// 0b111 == 7
	p(strconv.ParseInt("0b111", 0, 0))
	// 0111 == 73
	// 0 or 0o 被视为 8 进制
	p(strconv.ParseInt("0111", 0, 0))
	p(strconv.ParseInt("0o111", 0, 0))
	p(strconv.Atoi("1234"))
}
