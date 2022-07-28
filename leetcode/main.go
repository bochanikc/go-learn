package main

import "fmt"

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	fmt.Println(average(salary))

	fmt.Println(countOdds(3, 7))

}

// You are given an array of unique integers salary where salary[i] is the salary of the ith employee.
// Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.
func average(salary []int) float64 {
	min := salary[0]
	max := salary[0]
	sum := 0
	for _, value := range salary {
		//fmt.Println(value)
		if min > value {
			min = value
		}
		if max < value {
			max = value
		}
		sum += value

	}
	sum = sum - (min + max)
	//fmt.Println(min, max)
	avg := float64(sum) / float64(len(salary)-2)
	return avg
}

// Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).
func countOdds(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		//fmt.Println(i)
		if i%2 != 0 {
			count++
		}
	}
	return count
}
