package main

import (
	"fmt"
	"os"
)

func main() {

	type point struct {
		x, y int
	}
	p := point{4, 10}
	// {4, 10}
	fmt.Printf("struct1: %v\n", p)
	// {x:4 y:10}
	fmt.Printf("struct2: %+v\n", p)
	// main.point{x:4, y:10}
	fmt.Printf("struct3: %#v\n", p)
	// main.point
	fmt.Printf("type: %T\n", p)
	// %t bool
	fmt.Printf("bool: %t\n", 1 != 0)
	// 123
	fmt.Printf("int: %d\n", 123)
	// 1111011
	fmt.Printf("int: %b\n", 123)
	// 7b
	fmt.Printf("int: %x\n", 123)
	// {
	fmt.Printf("int: %c\n", 123)
	// 26.500000
	fmt.Printf("float1: %f\n", 26.5)
	// Hello,"世界"
	fmt.Printf("str1: %s\n", "Hello,\"世界\"")
	//  "Hello,\"世界\""
	fmt.Printf("str1: %q\n", "Hello,\"世界\"")
	// 字符串 hex
	// 48656c6c6f2c2022e4b896e7958c22
	fmt.Printf("str: %x\n", "Hello, \"世界\"")
	// point: 0x1400010c010
	fmt.Printf("point: %p\n", &p)
	// 格式化
	// width: |  1.20|      3.40|
	fmt.Printf("width: |%6.2f|%10.2f|\n", 1.2, 3.4)
	// width: |1.20  |3.40      |
	fmt.Printf("width: |%-6.2f|%-10.2f|\n", 1.2, 3.4)
	// width: |fo    |zz        |
	fmt.Printf("width: |%-6.2s|%-10.2s|\n", "foo", "zzf")
	// sprintf; Hello, World
	sprintf := fmt.Sprintf("sprintf; %s", "Hello, World")
	fmt.Println(sprintf)
	fmt.Fprintf(os.Stderr, "%10s", "Hello, World")

}
