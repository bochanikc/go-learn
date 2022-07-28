package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title  string
	Author string
}

// Объявляем тип Book, который удовлетворяет интерфейсу fmt.Stringer.
func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Объявляем тип Count, который удовлетворяет интерфейсу fmt.Stringer.
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Объявляем функцию WriteLog(), которая берёт любой объект,
// удовлетворяющий интерфейсу fmt.Stringer в виде параметра.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

type Battery struct {
	full string
}

func (b Battery) String() string {
	count := strings.Count(b.full, "1")
	strings.Repeat("подстрока_вставки", 2)
	newString := "[" + strings.Repeat(" ", 10-count) + strings.Repeat("X", count) + "]"
	return fmt.Sprintf("%s", newString)
}

func main() {
	batteryForTest := Battery{"1000010011"}
	WriteLog(batteryForTest)

	// Инициализируем объект Book и передаем в WriteLog().
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	count := Count(3)
	WriteLog(count)

	var i interface{} = 12

	if value, ok := i.(int); ok {
		fmt.Println(value + 12)
	}

	// Неизвестный интерфейс
	do(5)
	do("test")
	do(true)

	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции все полученные значения имеют тип пустого интерфейса
	checkValue(value1)
	checkValue(value2)
	doOperation(value1, value2, operation)

}

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.5, "*" //тут играемся с параметрами
}

func checkValue(i interface{}) {
	switch v := i.(type) {
	case float64:

	default:
		fmt.Printf("value=%v: %T \n", v, v)
		os.Exit(0)
	}
}

func doOperation(value1, value2, operation interface{}) {
	switch operation {
	case "+":
		result := value1.(float64) + value2.(float64)
		fmt.Printf("%.4f", result)
	case "-":
		result := value1.(float64) - value2.(float64)
		fmt.Printf("%.4f", result)
	case "*":
		result := value1.(float64) * value2.(float64)
		fmt.Printf("%.4f", result)
	case "/":
		result := value1.(float64) / value2.(float64)
		fmt.Printf("%.4f", result)
	default:
		fmt.Printf("неизвестная операция \n")
		os.Exit(0)
	}
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Умножим на 2:", v*2)
	case string:
		fmt.Println(v)
	default:
		fmt.Printf("Не знаю такого типа %T \n", v)
	}
}
