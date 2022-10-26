package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 0}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range kvs {
		fmt.Println("key: ", key)
	}

	fmt.Println("range для строк перебирает кодовые точки Unicode. Первое значение - это начальный байтовый индекс руны, а второе - сама руна.")
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
