package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Считывание ввода пользователя
	var input string
	fmt.Scanln(&input)

	// Разбиение ввода на операнды и операцию
	inputArr := strings.Split(input, " ") // сплит разделяет, мы как бы говорим что если будет пробел то разделить полученное значение
	if len(inputArr) != 3 {
		fmt.Println("Ошибка: введите два числа и операцию через пробел")
		return
	}

	// Парсинг операндов
	a, errA := parseOperand(inputArr[0])
	b, errB := parseOperand(inputArr[2])
	if errA != nil || errB != nil {
		fmt.Println("Ошибка: введите правильные числа")
		return
	}

	// Выполнение операции
	var result int
	switch inputArr[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неправильная операция")
		return
	}

	// выводим результат
	fmt.Println(result)
}

func parseOperand(s string) (int, error) {
	// проверяем на римлян
	if strings.ContainsAny(s, "IVXLCDM") {
		return parseRomanNumeral(s)
	}

	// парсинг арабскиз чисел
	num, err := strconv.Atoi(s)
	if err != nil || num < 1 || num > 10 {
		return 0, err
	}
	return num, nil
}

var romanNumeralvalues = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func parseRomanNumeral(s string) (int, error) {
	var result int
	var lastValue int

	for _, c := range s {
		value, ok := romanNumeralvalues[c]
		if !ok {
			return 0, fmt.Errorf("неправильно набраны римские символы: %c", c)
		}
		if value > lastValue && lastValue != 0 {
			result -= 2 * lastValue
		}
		result += value
		lastValue = value
	}
	if !isValidRomanNumeral(s, result) {
		return 0, fmt.Errorf("неправильное римское число: %s", s)
	}
	return result, nil
}

func isValidRomanNumeral(s string, value int) bool {
	if value <= 0 || value > 3999 {
		return false
	}
	// проверим что полученное значение соответствует исходной строке
	return toRomanNumeral(value) == s
}

var romanNumeralLetters = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func toRomanNumeral(n int) string {
	var result strings.Builder

	for _, numeral := range romanNumeralLetters {
		for n >= numeral.Value {
			result.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return result.String()
}
