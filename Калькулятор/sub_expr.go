package main
import(
	"strings"
)

func evalSubExpr(expr string) (float64, error) {
// Разделения выражения в скобках, прошедшего оценку на числа и операторы
	tokens := strings.Split(expr, "")
	nums := make([]float64, 0)
	ops := make([]rune, 0)
	numStr := ""
	negative := false

	err_float, tokens, nums, ops, numStr, negative := population(tokens,nums,ops,numStr,negative)
	if(err_float!=nil){
		return 0, err_float
	}
	// Расчет произведения и деления в первую очередь
	err, ops, nums := prod_dev(ops,nums)
	if(err!=nil){
		return 0, err
	}

	// Расчет суммирования и вычитания
	return sum_sub(ops,nums)
}