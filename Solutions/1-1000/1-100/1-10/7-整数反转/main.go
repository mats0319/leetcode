package mario

func reverse(x int) int {
	const (
		MAX = 2<<30 - 1
		MIN = -2 << 30
	)

	var result int

	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	if result > MAX || result < MIN {
		result = 0
	}

	return result
}
