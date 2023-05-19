package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	encoded1 := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("`%s` base64-> %v\n", data, encoded1)
	encoded2 := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Printf("`%s` base64-> %v\n", data, encoded2)
	decodeString, _ := base64.StdEncoding.DecodeString(encoded1)
	fmt.Println(string(decodeString))
	bytes, _ := base64.URLEncoding.DecodeString(encoded2)
	fmt.Println(string(bytes))
}
