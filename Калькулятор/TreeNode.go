package main

//описывем структуру узла бинарного дерва
type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
 }

func buildTree(tokens []string) *TreeNode {
	//Функция построения бинарного дерева. На вход принимает список 
	//обработанных элементов из строки вырожения
	//Объявляем два стека: ops (операторы) и nums (числа), 
	//которые будут содержит элементы, которые являются указателями на 
	//структуру TreeNode
	var ops []*TreeNode
	var nums []*TreeNode

	processOp := func(ops *[]*TreeNode, nums *[]*TreeNode) {
	//получаем значение последнего оператора из ops
	 op := (*ops)[len(*ops)-1]
	// удаляем значение последнего оператора из ops
	 *ops = (*ops)[:len(*ops)-1]
	 //получаем два последних значения числа из nums и присваиваем их соответственно
	 //переменным right и left. Удаляем их из исходного массива
	 right := (*nums)[len(*nums)-1]
	 *nums = (*nums)[:len(*nums)-1]
	 left := (*nums)[len(*nums)-1]
	 *nums = (*nums)[:len(*nums)-1]
	 //для оператора, который стал узлом устанавливаем числа в качестве потомков
	 op.Left = left
	 op.Right = right
	 //добавляем узел в nums
	 *nums = append(*nums, op)
	}
	pushOp := func(op *TreeNode) {
		//Запускаем цикл for, который продолжается до тех пор, 
		//пока длина среза ops больше нуля и приоритет операции 
		//op.Value меньше или равен приоритету операции последнего 
		//элемента в ops.
	 for len(ops) > 0 && precedence_operation(op.Value) <= precedence_operation(ops[len(ops)-1].Value) {
		processOp(&ops, &nums)
	 }
	 ops = append(ops, op)
	}
 
	for i := 0; i < len(tokens); i++ {
	 token := tokens[i]
	 //обрабатываем все случаи элементов из обработанной строки
	 switch token {
	 case "(":
		ops = append(ops, &TreeNode{Value: token})
	 case ")":
		for len(ops) > 0 && ops[len(ops)-1].Value != "(" {
		 processOp(&ops, &nums)
		}
		ops = ops[:len(ops)-1]
	 case "+", "-", "*", "/":
		pushOp(&TreeNode{Value: token})
		//если числовое значение просто добавляем в nums
	 default:
		nums = append(nums, &TreeNode{Value: token})
	 }
	}
 
	for len(ops) > 0 {
	 processOp(&ops, &nums)
	}
//возвращаем указатель на корень дерева
	return nums[0]
 }