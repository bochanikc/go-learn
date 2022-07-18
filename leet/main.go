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

	num := 234
	for _, value := range string(num) {
		fmt.Println(value)
	}
	// разбор числа по его числам
	n := 435465342634

	for {
		num := n % 10
		fmt.Println("num: ", num)
		n = n / 10
		fmt.Println(n)
		if n == 0 {
			break
		}
	}

	var inputBigNumber, inputSingleNumber, outputNumber string
	fmt.Scan(&inputBigNumber, &inputSingleNumber)

	for _, value := range inputBigNumber {
		if string(value) != inputSingleNumber {
			outputNumber += string(value)
		}
	}
	fmt.Println(outputNumber)

}
