package mario

func containsPattern(arr []int, m int, k int) bool {
	if len(arr) < m*k {
		return false
	}

	hasPattern := false
	for i := 0; i <= len(arr)-m*k; i++ {
		if isPattern(arr[i:i+m*k], m) {
			hasPattern = true
			break
		}
	}

	return hasPattern
}

func isPattern(slice []int, m int) bool {
	res := true
	for i := range slice {
		if slice[i] != slice[i%m] {
			res = false
			break
		}
	}

	return res
}
