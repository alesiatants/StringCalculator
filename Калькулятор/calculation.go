package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EvalMathExpr(expr string) (float64, error) {

	//применение функции валидации строки выражения
	if !IsValidExpression(expr) {
		return 0, fmt.Errorf("некорректное выражение")
	}

	// Оценка выражения в скобках
	for {
		start := strings.Index(expr, "(")
		if start == -1 {
			break
		}
		end := strings.Index(expr, ")")
		if end == -1 {
			return 0, fmt.Errorf("некорректное выражение")
		}
		subExpr := expr[start+1 : end]
		//применяем ркурсию для работы с подвыражениями
		subResult, err := EvalMathExpr(subExpr)

		if err != nil {
			return 0, err
		}

		expr = expr[:start] + strconv.FormatFloat(subResult, 'f', 6, 64) + expr[end+1:]
	}

	// Разделения выражения в скобках, прошедшего оценку на числа и операторы
	tokens := strings.Split(expr, "")
	nums := make([]float64, 0)
	ops := make([]rune, 0)
	numStr := ""
	negative := false

	for _, token := range tokens {
		//Пропускаем пробелы
		if token == " " {
			continue
		}
		/*Если символ является оператором +, * или /, она добавляет текущее число (если оно есть) в срез nums
		и сбрасывает переменную numStr. Она также добавляет оператор в срез ops.
		Если символ является оператором -, она проверяет, пуст ли numStr.
		Если да, она устанавливает флаг negative в true. Если нет, она добавляет текущее число в срез nums и
		сбрасывает переменную numStr. Добавляет оператор - в срез ops.
		Если символ является цифрой, она добавляет его в переменную numStr.*/
		if token == "+" || token == "*" || token == "/" {
			if numStr != "" {
				num, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					return 0, err
				}
				if negative {
					num = -num
					negative = false
				}
				nums = append(nums, num)
				numStr = ""
			}
			ops = append(ops, rune(token[0]))

		} else if token == "-" {

			if numStr == "" {
				negative = true
			} else {
				num, err := strconv.ParseFloat(numStr, 64)

				if err != nil {
					return 0, err
				}

				if negative {
					num = -num
					negative = false
				}

				nums = append(nums, num)
				numStr = ""
				ops = append(ops, '-')
			}

		} else {
			numStr += token
		}
	}

	if numStr != "" {
		num, err := strconv.ParseFloat(numStr, 64)

		if err != nil {
			return 0, err
		}

		if negative {
			num = -num
			negative = false
		}

		nums = append(nums, num)

	}

	// Расчет произведения и деления в первую очередь
	for i := 0; i < len(ops); i++ {

		if ops[i] == '*' || ops[i] == '/' {

			if ops[i] == '*' {
				nums[i] *= nums[i+1]
			} else {

				if nums[i+1] == 0 {
					return 0, fmt.Errorf("Деление на ноль!!!")
				}

				nums[i] /= nums[i+1]
			}
			nums = append(nums[:i+1], nums[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
			i--

		}
	}

	// Расчет суммирования и вычитания
	result := nums[0]

	for i := 0; i < len(ops); i++ {
		if ops[i] == '+' {
			result += nums[i+1]
		} else {
			result -= nums[i+1]
		}
	}

	return result, nil
}
