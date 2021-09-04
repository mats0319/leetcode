package mario

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	return calcPow(x, n)
}

func calcPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	quickMul := calcPow(x, n/2)

	res := quickMul * quickMul
	if n&1 == 1 {
		res *= x
	}

	return res
}
