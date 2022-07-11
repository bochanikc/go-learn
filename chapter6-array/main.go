package main

import "fmt"

func main() {
	// Объявление массива
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// Пример - оценка за экзамен
	mark := [5]float64{98, 93, 56, 87, 77}
	// var mark [5]float64
	// mark[0] = 98
	// mark[1] = 93
	// mark[2] = 56
	// mark[3] = 87
	// mark[4] = 77

	var total float64 = 0
	// for i := 0; i < len(mark); i++ {
	// 	total += mark[i]
	// }
	for _, value := range mark {
		total += value
	}
	fmt.Println(total / float64(len(mark)))

	// СРЕЗЫ slice
	var x1 []float64
	x2 := make([]float64, 5)
	x3 := make([]int, 5, 10)

	arr := [5]float64{1, 2, 3, 4, 5}
	x4 := arr[0:5]

	fmt.Println(x1, x2, x3, x4)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 10)
	copy(slice3, slice1)
	fmt.Println(slice3)

	// КАРТА map
	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"])

	m1 := make(map[int]int)
	m1[5] = 10
	fmt.Println(m1[5])
	delete(m1, 5)
	fmt.Println(m1[5])

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// Как обратиться к четвертому элементу массива или среза? - arr[4]

	// Чему равна длина среза, созданного таким способом: make([]int, 3, 9)? - 3
	fmt.Println(len(make([]int, 3, 9)))

	// Дан массив:
	arr1 := [6]string{"a", "b", "c", "d", "e", "f"} //что вернет вам x[2:5]? - "c","d","e"
	fmt.Println(arr1[2:5])

	// Напишите программу, которая находит самый наименьший элемент в этом списке:
	arr2 := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := arr2[0]
	for _, value := range arr2 {
		if value < min {
			min = value
		}
	}
	fmt.Println(min)

}
