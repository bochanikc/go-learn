package main

import "fmt"

func main() {
	var x string = "Hello world"
	fmt.Println(x)

	var x2 string
	x2 = "Hello world 2"
	fmt.Println(x2)

	x3 := 3
	fmt.Println(x3)

	catName := "Tigryzka"
	fmt.Println("My cat's name is", catName)

	const CONSTANTA string = "test"
	fmt.Println(CONSTANTA)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)

	// Farecheit to Celsi
	fmt.Println("Enter a Farecheit: ")
	var input int64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9
	fmt.Println("Celsius:", output)

}
