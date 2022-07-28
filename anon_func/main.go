package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T \n", x)

	fn := func(digit uint) uint {
		var newDigit uint
		newDigitString := strconv.FormatUint(uint64(digit), 10)
		var resultNewDigitString string
		for _, value := range newDigitString {
			valueInt, _ := strconv.Atoi(string(value))
			fmt.Println(valueInt)
			if valueInt%2 != 0 || valueInt == 0 {
				continue
			} else {
				resultNewDigitString += string(value)
			}
		}
		fmt.Println(resultNewDigitString)
		valueInt, _ := strconv.Atoi(resultNewDigitString)
		newDigit = uint(valueInt)
		if newDigit == 0 {
			return 100
		}
		return newDigit
	}

	fmt.Println(fn(30))
	ExampleDefer1()
}

func ExampleDefer1() {
	defer func() { fmt.Println(1) }()

	defer func() { fmt.Println(2) }()

	defer func() { fmt.Println(3) }()

	// Output:
	// 3
	// 2
	// 1
}
