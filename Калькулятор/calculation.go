package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EvalMathExpr(expr string) (float64, error) {
	//применение функции валидации строки выражения
	if !IsValidExpression(expr) {
		return 0, fmt.Errorf("Некорректное выражение")
	}
	// Оценка выражения в скобках
	// Удаление всех пробелов из выражения
	expr = strings.ReplaceAll(expr, " ", "")

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
		// Используем цикл for вместо рекурсии для оценки подвыражений
		subResult, err := evalSubExpr(subExpr)
		if err != nil {
			return 0, err
		}

		expr = expr[:start] + strconv.FormatFloat(subResult, 'f', 6, 64) + expr[end+1:]
	}
return evalSubExpr(expr)

}
