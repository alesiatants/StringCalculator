package main
import (
	"strings"
	"regexp"

)
// функция для проверки корректности выражения
func IsValidExpression(expression string) bool {
	// проверка на наличие только цифр, знаков базовых математических операций и круглых скобок
	if !regexp.MustCompile(`^[0-9\+\-\*\/\(\)\. ]+$`).MatchString(expression) {
		return false
	}
	
	expressionNoSpaces := strings.ReplaceAll(expression, " ", "")
	// проверка на содержание цифр перед закрывающимися скобками
	if strings.ContainsRune(expressionNoSpaces, 41){
		
		if !regexp.MustCompile(`\d\)`).MatchString(expressionNoSpaces) {
			return false
		}
	}
	
	// проверка на правильное использование круглых скобок (по количеству)
	if !IsValidBrackets(expression) {
		return false
	}

	return true
}

// функция для проверки правильного использования круглых скобок
func IsValidBrackets(expression string) bool {
	// подсчет количества открывающих и закрывающих скобок
	openBrackets := strings.Count(expression, "(")
	closeBrackets := strings.Count(expression, ")")

	// проверка на равенство количества открывающих и закрывающих скобок
	if openBrackets != closeBrackets {
		return false
	}
	//проверка на корректность следования скобок
		stack := make([]rune, 0)
		for _, c := range expression {
			switch c {
			case '(':
				stack = append(stack, c)
			case ')':
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
		
	return true
}

