package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else if i == 9 {
			fmt.Println("its ", i)
		}
	}

	a, b := 5, 4
	condition := "+"
	switch condition {
	case "+":
		fmt.Println(a, condition, b, "=", a+b)
	case "-":
		fmt.Println(a, condition, b, "=", a-b)
	default:
		fmt.Println("no this operation")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println(t, "It's before noon")
	default:
		fmt.Println(t, "It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know this type: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("stringa")
}
