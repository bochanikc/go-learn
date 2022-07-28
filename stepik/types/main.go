package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	a := 101.0 / 2.0
	fmt.Println(a + float64(uint8(a)))

	// int -> string
	b := strconv.Itoa(2020)
	fmt.Printf("%T \n", b)
	fmt.Println(b)

	user := "ученик"
	steps := 4
	fmt.Println("Поздравляю, " + user + "! Ты получил " + strconv.Itoa(steps) + "!")

	var a1 float64 = 1.0123456789
	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - точность (кол-во знаков после запятой)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	fmt.Println(strconv.FormatFloat(a1, 'f', 2, 64)) // 1.01

	// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
	fmt.Println(strconv.FormatFloat(a1, 'f', -1, 64)) // 1.0123456789

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	var b1 float64 = 2222 * 1023 * 245 * 2 * 52
	fmt.Println(strconv.FormatFloat(b1, 'e', -1, 64)) // 5.791874088e+10

	// bool -> string
	var a2 = true
	res := strconv.FormatBool(a2)
	fmt.Println(res)
	fmt.Printf("%T \n", res)

	//конвертация строк в целые числа
	// string->int
	foo := "10"
	bar := "15"
	fooInt, err := strconv.Atoi(foo)
	if err != nil {
		panic(err)
	}
	barInt, err := strconv.Atoi(bar)
	if err != nil {
		panic(err)
	}
	baz := fooInt - barInt
	fmt.Println(baz)

	// string->float
	s := "23.2342342"
	result, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("%f %T \n", result, result)

	//Конвертация string в bool
	s1 := "true"
	res1, err := strconv.ParseBool(s1)
	if err != nil { // не забываем проверить ошибку
		panic(err)
	}
	fmt.Println(res1)         // true
	fmt.Printf("%T \n", res1) // bool

	//Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.
	// Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки.
	// Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes.
	inputString1 := "%^80"
	inputString2 := "hhhhh20&&&&nd"

	resultinho := adding(inputString1, inputString2)
	fmt.Println(resultinho)

	// На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе
	// с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).
	inputStringCSV := "1 149,6088607594936;1 179,0666666666666"
	//var inputStringCSV string
	var resultAddCSV float64
	var digits []float64

	//inputStringCSV, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	splitString := strings.Split(inputStringCSV, ";")

	fmt.Println(splitString)

	for _, value := range splitString {
		digits = append(digits, changeStringNumber(value))
	}
	fmt.Println(digits[0], digits[1])
	resultAddCSV = digits[0] / digits[1]
	fmt.Printf("%.2f", resultAddCSV)

}

func changeStringNumber(stringNumber string) float64 {
	newStringNumber := strings.ReplaceAll(string(stringNumber), " ", "")
	newStringNumber = strings.ReplaceAll(string(newStringNumber), ",", ".")
	newNumber, _ := strconv.ParseFloat(newStringNumber, 64)
	return newNumber
}

func adding(inputString1 string, inputString2 string) int64 {
	inputString1, inputString2 = cleanString(inputString1), cleanString(inputString2)
	inputint1, _ := strconv.Atoi(inputString1)
	inputInt2, _ := strconv.Atoi(inputString2)
	result := inputint1 + inputInt2
	return int64(result)
}

func cleanString(inputString string) string {
	var newString string
	for _, value := range inputString {
		if unicode.IsDigit(value) {
			newString = newString + string(value)
		}
	}

	return newString
}
