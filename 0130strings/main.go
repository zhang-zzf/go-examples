package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 常量的声明和变量的声明不一样
	const s = "你好，世界"
	// strings are equivalent to []byte
	// string 底部是 []byte
	// len(s) = 15
	fmt.Println(s, len(s))

	for i := 0; i < len(s); i++ {
		// uint8 [0-255] 无符号 byte
		var u uint8 = s[i]
		// e4 bd a0 e5 a5 bd ef bc 8c e4 b8 96 e7 95 8c
		fmt.Printf("%x ", u)
	}

	fmt.Println()
	// 有 5 个字符（Rune）
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// 遍历字符串中的所有字符
	for idx, runeValue := range s {
		fmt.Printf("range-> %#U starts at %d\n", runeValue, idx)
	}

	// 遍历字符串中的所有字符
	for i, w := 0, 0; i < len(s); i += w {
		runeVal, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("utf8.DecodeRuneInString-> %#U starts at %d\n", runeVal, i)
		w = width
		checkRune(runeVal)
	}

}

func checkRune(r rune) {
	if r == '世' {
		fmt.Println("found 世")
	}
}
