package main
import(
	"fmt"
)

func prod_dev(ops []rune, nums []float64) (error, []rune, []float64) {
	for i := 0; i < len(ops); i++ {

		if ops[i] == '*' || ops[i] == '/' {

			if ops[i] == '*' {
				nums[i] *= nums[i+1]
			} else {

				if nums[i+1] == 0 {
					return fmt.Errorf("Деление на ноль!!!"), ops, nums
				}

				nums[i] /= nums[i+1]
			}
			nums = append(nums[:i+1], nums[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
			i--
		}
	}
return nil,ops,nums
}
