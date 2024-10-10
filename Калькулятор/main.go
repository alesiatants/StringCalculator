package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
	"io"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	file, err_open := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	for {

		fmt.Println(`
    Выберите действие

    1. Ввести строковое выражение
    2. Выйти

    `)
		fmt.Print(">>> ")

		choiceStr, err_read := reader.ReadString('\n')
		if err_read != nil {
			mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
			fmt.Println("Ошибка при чтении ввода:", err_read)
			continue // Продолжаем цикл, если возникла ошибка
		}
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)

		if err != nil {
		
			if err_open == nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				
				// перенаправление вывода логов в файл logfile.log
				log.Println("Неверный ввод! Введите числовой пункт меню.")
			} else {
				log.Panic("Ошибка открытия файла логов:", err_open)
			}
			defer file.Close() // отложенное закрытие файла
		
			continue

		}
		switch choice{
		case 1:
			fmt.Print("Введите строковое выражение для вычисления: \n>>> ")
			expressionStr, err_read := reader.ReadString('\n')
			if err_read != nil {
				mw := io.MultiWriter(os.Stdout, file)
					log.SetOutput(mw)
				fmt.Println("Ошибка при чтении ввода:", err_read)
				continue // Продолжаем цикл, если возникла ошибка
			}
			expressionStr = strings.TrimSpace(expressionStr)
			
			if IsValidExpression(expressionStr){
				root := parseExpression(expressionStr)
 				result, err := calculate(root)
				if err != nil {
				mw := io.MultiWriter(os.Stdout, file)
					log.SetOutput(mw)
				log.Println("Ошибка при вычислении выражения: ", err)
				continue
				} else {
					fmt.Printf("Результат вычисления выражения %s: %f\n", expressionStr, result)
				}
				
			} else {
				mw := io.MultiWriter(os.Stdout, file)
					log.SetOutput(mw)
				log.Println("Неккоректное выражение! ")
				continue
			}
		case 2:
			fmt.Println("Выход...")
			os.Exit(0)

		default:
			mw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(mw)
			log.Println("Неверный выбор. Введите корректный пункт меню (1 или 2).")
		}
	}
 }

func parseExpression(expression string) *TreeNode {
 return buildTree(tokenize(expression))
}

