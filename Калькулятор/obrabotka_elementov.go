package main

func tokenize(expression string) []string {
	//функция перебирает символы в строковом выражении и выделяет мат 
	//операции и числа разного состава
	var tokens []string
	for i := 0; i < len(expression); {
	 switch expression[i] {
	 case '(', ')', '+', '-', '*', '/':
		tokens = append(tokens, string(expression[i]))
		i++
	 default:
		start := i
		for i < len(expression) && (isDigit(expression[i]) || expression[i] == '.') {
		 i++
		}
		tokens = append(tokens, expression[start:i])
	 }
	}
	return tokens
 }

 func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
 }