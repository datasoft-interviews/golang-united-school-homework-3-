package homework

func reverse(input []int64) (result []int64) {
	res := make([]int64, len(input))
	copy(res, input)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
