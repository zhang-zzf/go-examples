package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash := sha256.New()
	s := "sha256 a string"
	hash.Write([]byte(s))
	bs := hash.Sum(nil)
	// 32 bytes
	fmt.Printf("%v hash code -> 0x%x\n", s, bs)

	h := sha256.New()
	h.Write([]byte("sha256 "))
	h.Write([]byte("a string"))
	fmt.Printf("%v hash code -> 0x%x\n", s, h.Sum(nil))
}
