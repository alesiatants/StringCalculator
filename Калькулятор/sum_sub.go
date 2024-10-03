package main

func sum_sub(ops []rune, nums []float64) (float64, error) {
	result := nums[0]
	for i := 0; i < len(ops); i++ {
		if ops[i] == '+' {
			result += nums[i+1]
		} else {
			result -= nums[i+1]
		}
	}
	return result, nil
}