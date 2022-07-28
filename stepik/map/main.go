package main

import (
	"fmt"
)

func main() {
	users1 := make(map[int]string)

	users2 := map[int]string{
		1:  "Borow",
		10: "Promes",
	}

	fmt.Println(users1)
	fmt.Println(users2)

	fmt.Println(users2[10])

	delete(users2, 1)
	fmt.Println(users2[1])

	// Проверка на пустое значение
	m := map[int]int{
		1: 10,
	}
	if value, inMap := m[1]; inMap {
		fmt.Println(value)
	}
	if value, inMap := m[2]; inMap {
		fmt.Println(value) // Условие не выполняется
	}

	// len()
	m1 := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	fmt.Println(len(m1))

	// база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation в т.ч.
	// была сохранена информация о городах, которые входят в другие группы из groupCity.
	// Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation,
	// чтобы в ней была сохранена информация только о городах из группы groupCity[100].
	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село"},        // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионник"},             // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Село":          100,
		"Город":         300,
		"Большой город": 700,
		"Миллионник":    500,
	}

	fmt.Println(groupCity[100][1])
	fmt.Println(cityPopulation["Село"])

	var flag bool
	for key, _ := range cityPopulation {
		fmt.Println("cityPopulationKey:", key)
		flag = true
		for _, valueGP := range groupCity[100] {
			fmt.Println("valueGP:", valueGP)
			if key == valueGP {
				flag = true
				break
			} else {
				flag = false
			}
		}
		if !flag {
			fmt.Println("delete: ", key, cityPopulation)
			delete(cityPopulation, key)
		}
	}
	fmt.Println(groupCity)
	fmt.Println(cityPopulation)
}
