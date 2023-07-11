package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {

	fmt.Println("1+2=", plus(1, 2))
	fmt.Println("1+2+3=", plus2(1, 2, 3))

	a, b := values()
	fmt.Println("a,b=", a, b)

	sum1 := sum(1, 2, 3)
	fmt.Println("1+2+3=", sum1)

	nums := []int{1, 2, 3}
	// unpack slice
	fmt.Println("sum nums=", sum(nums...))

	funcA := intSeq()
	fmt.Println("funcA:", funcA(), funcA())
	funcB := intSeq()
	fmt.Println("funcA:", funcB(), funcB())

	fmt.Println("fact 5=", fact(5))

	var fib func(int) int
	fib = func(num int) int {
		if num == 1 || num == 2 {
			return 1
		}
		return fib(num-1) + fib(num-2)
	}
	fmt.Println("fib 5=", fib(5))

}

func sum(nums ...int) int {
	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}

func plus2(a, b, c int) int {
	return a + b + c
}

func values() (int, int) {
	return 1, 2
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}
