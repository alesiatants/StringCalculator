package main

import(
	"strconv"
	"fmt"
)

func precedence_operation(op string) int {
	//описание приоритета выполнения операций
	switch op {
	case "+", "-":
	 return 1
	case "*", "/":
	 return 2
	default:
	 return 0
	}
 }
 
 func calculate(node *TreeNode) (float64, error) {
	 if node == nil {
			 return 0, fmt.Errorf("Дерево не заполнено")
	 }
	 stack := []*TreeNode{}
	 var curr *TreeNode
	 values := map[*TreeNode]float64{}
	 visited := map[*TreeNode]bool{}
 
	 stack = append(stack, node)
 
	 for len(stack) > 0 {
			 curr = stack[len(stack)-1]
			 // Если это листовой узел (число)
			 if curr.Left == nil && curr.Right == nil {
					 val, err := strconv.ParseFloat(curr.Value, 64)
					 if err!=nil{
						return 0,err
					 }
					 values[curr] = val
					 stack = stack[:len(stack)-1]
					 continue
			 }
 
			 if visited[curr] {
					 leftVal := values[curr.Left]
					 rightVal := values[curr.Right]
					 var result float64
					 //производим вычисления операции-узла над его потомками
 
					 switch curr.Value {
					 case "+":
							 result = leftVal + rightVal
					 case "-":
							 result = leftVal - rightVal
					 case "*":
							 result = leftVal * rightVal
					 case "/":
							if rightVal==0{
								return 0, fmt.Errorf("Деление на ноль!")
							}
							 result = leftVal / rightVal
					 }
					 
					 values[curr] = result
					 stack = stack[:len(stack)-1]
			 } else {
					 if curr.Right != nil {
							 stack = append(stack, curr.Right)
					 }
					 if curr.Left != nil {
							 stack = append(stack, curr.Left)
					 }
					 visited[curr] = true
			 }
	 }
	 return values[node], nil
 }
 
