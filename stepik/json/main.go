package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// json.Marshal() - кодирование данных в json
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", data)

	// Человеческий вывод json
	// MarshalIndent похож на Marshal, но применяет отступ (indent) для форматирования вывода.
	//Каждый элемент JSON в выходных данных начинается с новой строки, начинающейся с префикса (prefix), за которым следует один или несколько отступов в соответствии с вложенностью
	data2, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", data2)

	// Unmarshall - json -> struct
	dataUnmarshall := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

	var sUnmarshall myStruct

	if err := json.Unmarshal(dataUnmarshall, &sUnmarshall); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", s)

	// в общем виде аннотация выглядит так: `json:"используемое_имя,опция,опция"`
	type myStruct1 struct {
		// при кодировании / декодировании будет использовано имя name, а не Name
		Name string `json:"name"`

		// при кодировании / декодировании будет использовано то же имя (Age),
		// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
		// то при кодировании оно будет опущено
		Age int `json:",omitempty"`

		// при кодировании / декодировании поле всегда игнорируется
		Status bool `json:"-"`
	}

	m := myStruct1{Name: "John Connor", Age: 0, Status: true}

	data3, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data3) // {"name":"John Connor"}

	//В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
	// Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
	// Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
	// {
	// 	"Average": 14.1
	// }

	jsonStudentsData1 := []byte(`{
		"ID":134,
		"Number":"ИЛМ-1274",
		"Year":2,
		"Students":[
			{
				"LastName":"Вещий",
				"FirstName":"Лифон",
				"MiddleName":"Вениаминович",
				"Birthday":"4апреля1970года",
				"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
				"Phone":"+7(948)709-47-24",
				"Rating":[1,2]
			},
			{
				"LastName":"Вещий2",
				"FirstName":"Лифон",
				"MiddleName":"Вениаминович",
				"Birthday":"4апреля1970года",
				"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
				"Phone":"+7(948)709-47-24",
				"Rating":[3]
			}
		]
	}`)
	// fmt.Println(jsonStudentsData1)
	// jsonStudentsData, err := ioutil.ReadAll(os.Stdin)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(getStudentsRating(jsonStudentsData1))
}

func getStudentsRating(jsonStudentsData []byte) string {
	type Students struct {
		LastName   string
		FirstName  string
		MiddleName string
		Birthday   string
		Address    string
		Phone      string
		Rating     []int
	}
	type StudentGroup struct {
		ID       int
		Number   string
		Year     int
		Students []Students
	}

	type Rating struct {
		Average float64
	}

	var testStudentGroup StudentGroup

	if err := json.Unmarshal(jsonStudentsData, &testStudentGroup); err != nil {
		fmt.Println(err)
		return ""
	}

	countStudent := 0
	summMarks := 0
	for _, value := range testStudentGroup.Students {
		//fmt.Println(value)
		// for _, valueMark := range value.Rating {
		// 	summMarks += 1
		// }
		summMarks += len(value.Rating)
		countStudent += 1
	}
	//fmt.Println(summMarks, countStudent)
	resultAverage := float64(summMarks) / float64(countStudent)
	//fmt.Println(resultAverage)
	resultAverageStruct := Rating{
		Average: resultAverage,
	}
	dataResultAverage, err := json.MarshalIndent(resultAverageStruct, "", "    ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("%s", dataResultAverage)
}
