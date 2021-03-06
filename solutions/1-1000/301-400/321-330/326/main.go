package mario

// An integer n is a power of three, if there exists an integer x such that n == 3^x
func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	isValid := false
	for mul, param := 1, 3; mul <= n; mul *= param {
		if mul == n {
			isValid = true
			break
		}
	}

	return isValid
}
