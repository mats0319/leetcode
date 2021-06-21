package mario

// temperatures.length >= 1
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 1) // index stack, stack[0] = 0 is right
	res := make([]int, len(temperatures))
	for i := 1; i < len(res); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			index := stack[len(stack)-1]
			res[index] = i - index

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
