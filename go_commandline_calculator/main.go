package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsRoman(roman string) bool {
	if roman == "" {
		return false
	}

	romanb := []byte(strings.ToUpper(roman))
	check, _ := regexp.Match("^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", romanb)
	return check
}

var RomanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := RomanNumerals[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("-----------------------------------------------------------------------------------------------------------")
		fmt.Println("Калькулятор умеет работать в режиме арабские-арабские или римские-римские, будьте внимательны.")
		fmt.Println("Вводите значения в формате: операнд1 операция операнд2. Иные форматы ввода будут расценены как ошибочные!")
		fmt.Println("Напоминание: в римской системе счисления не существует отрицательных чисел)")
		text, _ := reader.ReadString('\n')
		words := strings.Fields(text)

		if len(words) == 3 {
			operator := words[1]
			if IsRoman(words[0]) && IsRoman(words[2]) {
				operand1, operand2 := romanToInt(strings.ToUpper(words[0])), romanToInt(strings.ToUpper(words[2]))

				switch operator {
				case "+":
					fmt.Println(integerToRoman(operand1 + operand2))
					return
				case "-":
					result := operand1 - operand2
					if result < 1 {
						fmt.Println("Результатом работы калькулятора с римскими числами могут быть только положительные числа!")
						return
					} else {
						fmt.Println(integerToRoman(result))
						return
					}
				case "*":
					fmt.Println(integerToRoman(operand1 * operand2))
					return
				case "/":
					result := operand1 / operand2
					if result < 1 {
						fmt.Println("Результатом работы калькулятора с римскими числами могут быть только положительные числа!")
						return
					} else {
						fmt.Println(integerToRoman(result))
						return
					}
				default:
					fmt.Println("Ошибка, неправильный формат ввода!")
					return
				}
			} else {
				operand1, err1 := strconv.Atoi(words[0])
				operand2, err2 := strconv.Atoi(words[2])
				if err1 != nil || err2 != nil {
					fmt.Println("Ошибка, неправильный формат ввода!")
					return
				} else {
					switch operator {
					case "+":
						fmt.Println(operand1 + operand2)
						return
					case "-":
						fmt.Println(operand1 - operand2)
						return
					case "*":
						fmt.Println(operand1 * operand2)
						return
					case "/":
						fmt.Println(operand1 / operand2)
						return
					default:
						fmt.Println("Ошибка, неправильный формат ввода!")
						return
					}
				}
			}
		} else {
			fmt.Println("Ошибка, неправильный формат ввода!")
			return
		}
	}
}
