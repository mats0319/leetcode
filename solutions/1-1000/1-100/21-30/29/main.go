package mario

func divide(dividend int, divisor int) int {
	isNeg := 0
	if dividend < 0 {
		dividend *= -1
		isNeg++
	}
	if divisor < 0 {
		divisor *= -1
		isNeg++
	}

	res := calc(dividend, divisor)
	if isNeg&1 == 1 {
		res *= -1
	}

	if res < -1<<31 || res > 1<<31-1 {
		res = 1<<31 - 1
	}

	return res
}

func calc(dividend, divisor int) int {
	if dividend == divisor {
		return 1
	} else if dividend < divisor {
		return 0
	}

	times := 0
	for dividend > divisor {
		divisor <<= 1
		times++
	}

	return 1<<(times-1) + calc(dividend-divisor>>1, divisor>>times)
}
