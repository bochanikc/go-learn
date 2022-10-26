package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}
	for _, value := range numbers {
		sum += value
	}
	return sum
}
