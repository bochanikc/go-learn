package main

import "fmt"

func main() {
	fmt.Println("5 + 3 =", 5+3)
	fmt.Println("Kia"[1])
	fmt.Println(len("Kia"))
	fmt.Println("Kia " + "Pride")

	//bool
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
