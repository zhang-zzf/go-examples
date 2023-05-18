package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type resp1 struct {
	Age    int
	Fruits []string
}

type resp2 struct {
	Age    int      `json:"age"`
	Fruits []string `json:"fruits"`
}

func main() {
	bytes, _ := json.Marshal(1)
	fmt.Println(string(bytes))
	bytes, _ = json.Marshal(1.25)
	fmt.Println(string(bytes))
	bytes, _ = json.Marshal(true)
	fmt.Println(string(bytes))
	bytes, _ = json.Marshal(`"Hello, World"`)
	fmt.Println(string(bytes))
	bytes, _ = json.Marshal([]string{"Hello", "World"})
	fmt.Println(string(bytes))
	bytes, _ = json.Marshal(map[string]any{
		"key1": 1,
		"key2": true,
		"key3": 1.25,
		"key4": []string{"Hello", "World"},
		"key5": map[string]string{
			"m1": "m1Value",
		},
	})
	fmt.Println(string(bytes))

	r := &resp1{Age: 10, Fruits: []string{"apple", "orange"}}
	bytes, _ = json.Marshal(r)
	fmt.Println(string(bytes))

	rr1 := resp1{}
	json.Unmarshal(bytes, &rr1)
	fmt.Println(rr1)

	r2 := &resp2{Age: 10, Fruits: []string{"apple", "orange"}}
	bytes, _ = json.Marshal(r2)
	fmt.Println(string(bytes))

	rr2 := resp2{}
	json.Unmarshal(bytes, &rr2)
	fmt.Println(rr2)

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(rr2)

}
