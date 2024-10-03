package main
import (
	"strconv"
)
func population(tokens []string, nums []float64, ops []rune, numStr string, negative bool)(error,[]string,[]float64,[]rune,string,bool){
	for _, token := range tokens {
		//Пропускаем пробелы
			if token == " " {
				continue
			}

			if token == "+" || token == "*" || token == "/" {
				if numStr != "" {
					num, err := strconv.ParseFloat(numStr, 64)
					if err != nil {
						return err, tokens, nums, ops, numStr, negative
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
						return err, tokens, nums, ops, numStr, negative
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
				return err, tokens, nums, ops, numStr, negative
			}
	
			if negative {
				num = -num
				negative = false
			}
	
			nums = append(nums, num)
	
		}
		return nil, tokens, nums, ops, numStr, negative
	}