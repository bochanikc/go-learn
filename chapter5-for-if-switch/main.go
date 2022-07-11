package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, " even")
		} else if j%2 != 0 {
			fmt.Println(j, " odd")
		}
	}

	var key int64
	fmt.Println("Enter 0-3 case: ")
	fmt.Scanf("%d", &key)
	switch key {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Unknown number")
	}

}
