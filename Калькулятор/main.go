package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println(`
    Выберите действие

    1. Ввести строковое выражение
    2. Выйти

    `)
		fmt.Print(">>> ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)

		if err != nil {

			fmt.Println("Неверный ввод!")
			continue

		}

		if choice == 1 {

			fmt.Print("Введите строковое выражение для вычисления: \n>>> ")
			expressionStr, _ := reader.ReadString('\n')
			expressionStr = strings.TrimSpace(expressionStr)
			result, err := EvalMathExpr(expressionStr)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}

		} else if choice == 2 {

			fmt.Println("Выход...")
			os.Exit(0)

		} else {

			fmt.Println("Неверный выбор. Введите корректный пункт меню (1 или 2).")

		}
	}

}
