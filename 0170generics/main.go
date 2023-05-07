package main

import "fmt"

func mapKeys[K comparable, V any](m map[K]V) []K {
	var ret []K
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	val  T
	next *Node[T]
}

func (list *List[T]) Push(val T) {
	n := &Node[T]{val: val}
	if list.tail == nil {
		list.head = n
		list.tail = n
	} else {
		list.tail.next = n
		list.tail = n
	}
}

func (list *List[T]) Take() (T, bool) {
	if list.head == nil {
		// 定义 zero 值
		var zero T
		return zero, true
	} else {
		val := list.head.val
		list.head = list.head.next
		if list.head == nil {
			list.tail = nil
		}
		return val, false
	}
}

func main() {
	m := map[int]string{1: "1", 2: "2"}
	keys := mapKeys(m)
	fmt.Println("keys:", keys)

	// 指定范型
	_ = mapKeys[int, string](m)

	list := List[int]{}
	list.Push(1)
	// Cannot use '"String"' (type string) as the type T (int
	// list.Push("String")
	list.Push(2)

	fmt.Println(list.Take())
	fmt.Println(list.Take())
	fmt.Println(list.Take())
	fmt.Println(list.Take())
}
