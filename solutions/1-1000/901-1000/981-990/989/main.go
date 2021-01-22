package mario

// K >= 0
func addToArrayForm(A []int, K int) []int {
	// make opposite direction array-form K
	ks := make([]int, 0)
	{
		for K > 0 {
			ks = append(ks, K%10)
			K /= 10
		}
	}

	// revert A into opposite direction
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}

	// calc opposite direction result
	res := make([]int, max(len(A), len(ks))+1)
	{
		addFlag := 0
		for i := 0; i < len(res); i++ {
			digit := addFlag
			{
				if i < len(A) {
					digit += A[i]
				}
				if i < len(ks) {
					digit += ks[i]
				}
			}

			addFlag = digit / 10

			res[i] = digit % 10
		}
	}

	// revert result and trim meaningless leading zero
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	if len(res) > 1 && res[0] == 0 {
		res = res[1:]
	}

	return res
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
