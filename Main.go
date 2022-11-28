package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var RomanNumerals = map[rune]int{
	//Про руны нет информации в приложенных материалах, но их природа не сильно отличается от byte\char в C-подобных языках. По сути- конвертация символа в числовой u32 код.
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

func romanToInt(s string) int {
	//TODO: ошибка при некорректно сформулированных римских цифрах (IIIIIIIIX или XLX)
	Total := 0
	MaxOrder := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]

		num := RomanNumerals[rune(letter)]
		if num >= MaxOrder {
			MaxOrder = num
			Total += num
			continue
		}
		Total -= num
	}
	return Total
}
func ArabicToRoman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{ //нам достаточно вывода до 100, т.к. максимальное значение будет при операции 10*10.
		//В данном случае использованы структуры.
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

	RomanOutput := ""
	//Пустой идентификатор для обхода необходимости использования переменных. В данном случае- счётчика\итератора
	for _, conversion := range conversions {
		for number >= conversion.value {
			RomanOutput += conversion.digit
			number -= conversion.value
		}
	}
	return RomanOutput
}
func main() {
	//if-else по большей части инвертированы для более удобочитаемого кода.
	//Считываем ввод с консоли
	var Expression string
	//Используем bufio для считывания строки вместе с пробелами.
	Input := bufio.NewScanner(os.Stdin)
	Input.Scan()
	Expression = Input.Text()
	//Удаляем пробелы для облегчения обработки.
	Expression = strings.Replace(Expression, " ", "", -1)

	//проверяем наличие операндов

	var OperatorsCounter int = 0
	var result int = 0
	var RomanNumUsed bool = false
	var ArabicNumUsed bool = false
	var ExpectedOperator string = ""
	var Operators = [4]string{
		"+",
		"-",
		"*",
		"/",
	}
	if !(strings.ContainsAny(Expression, "+-*/")) {
		err := fmt.Errorf("В строке нет операторов!")
		fmt.Print(err)
	} else {
		for Key := range Operators {
			var CurrentOperatorCounter int = strings.Count(Expression, Operators[Key])
			OperatorsCounter += CurrentOperatorCounter
			if CurrentOperatorCounter > 0 {
				ExpectedOperator = Operators[Key]
			}
		}
		if OperatorsCounter != 1 {
			err := fmt.Errorf("More than 1 operator!")
			fmt.Print(err)
		} else {
			//У нас в строке один оператор, и он нам известен (Лежит в переменной ExpectedOperator). Проверим виды вводимых чисел- арабские\римские, исключим смешивание.
			if strings.ContainsAny(Expression, "0123456789") {
				ArabicNumUsed = true
			}
			if strings.ContainsAny(Expression, "IVX") {
				RomanNumUsed = true
			}
			for _, letter := range Expression {
				if !strings.Contains("012345689IiXxVvLlCcMm", string(letter)) {
					panic("В вводе обнаружены символы помимо римских или арабских цифр!")
				}
			}
			if RomanNumUsed && ArabicNumUsed {
				err := fmt.Errorf("Арабские и римские цифры перемешаны!")
				fmt.Print(err)
			} else {
				if !(RomanNumUsed || ArabicNumUsed) {
					err := fmt.Errorf("В вводе не найдено ни римских, ни арабских цифр!")
					fmt.Print(err)
				} else {
					res := strings.Split(Expression, ExpectedOperator)
					if res[1] == "" || res[0] == "" {
						err := fmt.Errorf("Оператор в начале или конце строки! Нет второго числа.")
						fmt.Print(err)
					} else {

						if ArabicNumUsed {
							result1, err := strconv.Atoi(res[0])
							result2, err := strconv.Atoi(res[1])
							if !(result1 <= 10 && result2 <= 10) {
								err := "Использованы числа выше 10!"
								fmt.Print(err)
							} else {
								if err != nil {
									panic(err)
								}
								switch ExpectedOperator {
								case "+":
									result = result1 + result2
								case "-":
									result = result1 - result2
								case "*":
									result = result1 * result2
								case "/":
									result = result1 / result2

								}
								fmt.Print(result)
							}
						} else {
							result1 := romanToInt(res[0])
							result2 := romanToInt(res[1])
							if result1 > 10 || result2 > 10 {
								err := fmt.Errorf("Использованы числа выше 10!")
								fmt.Print(err)
							} else {
								switch ExpectedOperator {
								case "+":
									result = result1 + result2
								case "-":
									if result1-result2 > 0 {
										result = result1 - result2
									} else {
										err := fmt.Errorf("Римские числа не могут быть отрицательными!")
										fmt.Print(err)
									}
								case "*":
									result = result1 * result2
								case "/":
									{
										result = result1 / result2
									}
								}
								RomanResult := ArabicToRoman(result)
								fmt.Print(RomanResult)
							}
						}

					}
				}
			}
		}
	}
}
