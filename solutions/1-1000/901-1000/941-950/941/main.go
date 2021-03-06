package mario

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	i := 1
	isValid := false
	for {
		for i < len(A) && A[i-1] < A[i] {
			i++
			isValid = true
		}
		if !isValid {
			break
		}

		isValid = false
		for i < len(A) && A[i-1] > A[i] {
			i++
			isValid = true
		}
		if !isValid {
			break
		}

		break
	}

	return isValid && i == len(A)
}
