package mario

func plusOne(digits []int) []int {
	addFlag := 1
	for i := len(digits) - 1; i >= 0 && addFlag > 0; i-- {
		value := addFlag + digits[i]

		addFlag = value / 10

		digits[i] = value % 10
	}

	if addFlag > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
