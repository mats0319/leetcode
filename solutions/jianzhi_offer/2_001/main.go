package mario

// divide '*', '/', '%' is forbidden
func divide(a int, b int) int {
	sign := 1
	if a < 0 {
		sign *= -1
		a *= -1
	}
	if b < 0 {
		sign *= -1
		b *= -1
	}

	res := 0

	for a >= b {
		rate := 1
		value := b
		for a >= value+value {
			rate += rate
			value += value
		}

		res += rate
		a -= value
	}

	res *= sign
	if res > 1<<31-1 {
		res = 1<<31 - 1
	}

	return res
}
