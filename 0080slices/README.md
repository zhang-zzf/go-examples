# Slices

## Array

An array type definition specifies a length and an element type

> `[4]int` and `[5]int` are distinct, incompatible types

- `var a1 [4]int`
- `a2 := [2]int{1,2}`
- `a3 := [...]int{1,2}`

## Slices

- `var s1 []int`
- `s2 := []int{1,2}`