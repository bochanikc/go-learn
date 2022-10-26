package main

import "fmt"

func main() {
	sum(1, 4, 5, 9)
	sum(0)
	sum(4, 4, 4)

	fmt.Println("Если у вас есть несколько аргументов в срезе, вы можете применить его к функции с переменным числом аргументов таким образом func(slice...)")
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	result := 0
	for _, value := range nums {
		result += value
	}
	fmt.Println(result)
}
