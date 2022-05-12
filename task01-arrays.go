package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum / float32(len(input))
}
