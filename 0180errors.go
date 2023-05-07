package main

import (
	"errors"
	"fmt"
)

func f1(i int) (int, error) {
	if i == 0 {
		return 0, errors.New("i is 0")
	}
	return 5 / i, nil
}

type argError struct {
	arg int
	msg string
}

func (r *argError) Error() string {
	return fmt.Sprintf("%d->%s", r.arg, r.msg)
}

func f2(i int) (int, error) {
	if i == 0 {
		return 0, &argError{i, "i is zero"}
	}
	return 5 / i, nil
}

func main() {

	fmt.Println(f1(0))
	fmt.Println(f2(0))

	for _, i := range []int{1, 2, 0} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	_, e := f2(0)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg, ae.msg)
	}
}
