package main

import (
	"fmt"
	"time"
)

func main() {

	// case 不用 break
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't known type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(5)
	whatAmI(whatAmI)

	str := ""
	switch str {
	case "":
		fmt.Println("")
	case "H":
		fmt.Println("H")
	default:
		fmt.Println("unKnown")
	}

}
