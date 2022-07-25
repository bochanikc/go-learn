package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	// На вход подается строка. Нужно определить, является ли она правильной или нет.
	// Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textRune := []rune(text)

	//fmt.Println(string(textRune[0]))
	if unicode.IsUpper(textRune[0]) && textRune[utf8.RuneCountInString(string(textRune))-2] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
	//fmt.Println(string(textRune[utf8.RuneCountInString(string(textRune))-2]))

	//На вход подается строка. Нужно определить, является ли она палиндромом.
	// Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
	// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").
	//text, _ = bufio.NewReader(os.Stdin).ReadString(' ')
	//text = text[:len(text)]
	fmt.Scan(&text)

	//fmt.Println(text)
	//fmt.Println(reverse(text))

	if text == reverse(text) {
		//fmt.Println("Палиндром")
	} else {
		//fmt.Println("Нет")
	}

	//  На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
	codeString := "ihgewlqlkot"
	decodeString := ""

	codeStringRune := []rune(codeString)

	for i := range codeStringRune {
		if i == 0 {
			continue
		} else if i%2 != 0 {
			decodeString = decodeString + string(codeStringRune[i])
		}
	}
	fmt.Print(string(decodeString))

	// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
	codeString = "zaabcbd"
	decodeString = ""

	codeStringRune = []rune(codeString)

	for _, value := range codeStringRune {
		fmt.Println(strings.Count(codeString, string(value)))
		if strings.Count(codeString, string(value)) > 1 {
			continue
		} else {
			decodeString = decodeString + string(value)
		}
	}

	fmt.Println(string(decodeString))

	// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
	checkPassword := "fdsghdfgjsdDD1"
	var flag bool

	checkPasswordRune := []rune(checkPassword)

	if len(checkPasswordRune) >= 5 {
		for _, value := range checkPasswordRune {
			if unicode.Is(unicode.Latin, value) || unicode.IsDigit(value) {
				flag = true
				continue
			} else {
				flag = false
				break
			}
		}

	} else {
		flag = false
	}

	if flag {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}

	// Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
	fmt.Println(findMaxDigit("1112221112"))
	fmt.Println(squareDigits(9119))
}

func squareDigits(inputInt int) int {
	newDigit := ""

	inputRune := []rune(strconv.Itoa(inputInt))

	for _, value := range inputRune {
		addValue, _ := strconv.Atoi(string(value))
		newDigit = newDigit + strconv.Itoa(addValue*addValue)
	}
	result, _ := strconv.Atoi(newDigit)
	return result
}

func findMaxDigit(inputString string) int {
	maxDigit := 0
	for i := range inputString {
		value, _ := strconv.Atoi(string(inputString[i]))
		if value > maxDigit {
			maxDigit = value
		}
	}
	return maxDigit
}

func reverse(s string) string {
	stringRune := []rune(s)
	count := len(stringRune)
	newStringRune := make([]rune, count)

	for _, c := range stringRune {
		count--
		newStringRune[count] = c
	}
	return string(newStringRune)
}
