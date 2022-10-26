package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 23

	fmt.Println("map - ", m)

	v1 := m["k1"]
	fmt.Println("v1 - ", v1)

	fmt.Println("len - ", len(m))

	delete(m, "k1")
	fmt.Println("map - ", m)

	_, keyIsExt := m["k2"]
	fmt.Println("key is exist in map? ", keyIsExt)

	n := map[string]int{"str1": 1, "str2": 2}
	fmt.Println("map2 - ", n)
}
