package main

import (
	"fmt"
	"strconv"
)

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	fmt.Println(average(salary))

	fmt.Println(countOdds(3, 7))

	fmt.Println(hammingWeight(00000000000000000000000000001011))

	fmt.Println(subtractProductAndSum(234))

	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1}))

	fmt.Println(runningSum([]int{1, 2, 3, 4, 5}))
}

// There is a function signFunc(x) that returns:
// 1 if x is positive.
// -1 if x is negative.
// 0 if x is equal to 0.
// You are given an integer array nums. Let product be the product of all values in the array nums.
// Return signFunc(product).
func arraySign(nums []int) int {
	mark := 1
	for _, value := range nums {
		if value == 0 {
			return 0
		} else if value < 0 {
			mark = -mark
		}
	}
	return mark
}

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//Return the running sum of nums.
func runningSum(nums []int) []int {
	result := []int{}
	count := 0
	for _, value := range nums {
		count += value
		result = append(result, count)
	}
	return result
}

func largestPerimeter(nums []int) int {

	fmt.Println(nums[0], nums[1], nums[2])
	if len(nums) == 3 && nums[0] < (nums[1]+nums[2]) && nums[1] < (nums[0]+nums[2]) && nums[2] < (nums[0]+nums[1]) {
		p := nums[0] + nums[1] + nums[2]
		return p
	} else {
		return 0
	}
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

// Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
func hammingWeight(num uint32) int {
	count := 0
	numString := []rune(strconv.FormatUint(uint64(num), 2))
	for _, value := range numString {
		//fmt.Println(string(value))
		if string(value) == "1" {
			count += 1
		}
	}
	return count
}

// Given an integer number n, return the difference between the product of its digits and the sum of its digits.
func subtractProductAndSum(n int) int {
	summ := 0
	product := 1
	numString := []rune(strconv.Itoa(n))
	for _, value := range numString {
		//fmt.Println(string(value))
		digit, _ := strconv.Atoi(string(value))
		summ += digit
		product *= digit
	}
	result := product - summ
	return result
}
