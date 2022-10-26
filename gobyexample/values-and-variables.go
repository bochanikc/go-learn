package main

import "fmt"

func main() {
	var summ = 1 + 1
	div := 7.0 / 3.0

	var (
		resTrue  = true
		resFalse = false
	)

	const n = 500000

	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", summ)
	fmt.Println("7.0/3.0 =", div)

	fmt.Println(resTrue)
	fmt.Println(resFalse)
	fmt.Println(!false)
	fmt.Println(n)
}
