package mario

func longestMountain(A []int) int {
	if len(A) < 3 {
		return 0
	}

	res := 0
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			left, right := i-1, i+1
			for extendFlag := true; extendFlag; {
				extendFlag = false
				if left >= 0 && A[left] < A[left+1] {
					left--
					extendFlag = true
				}
				if right < len(A) && A[right-1] > A[right] {
					right++
					extendFlag = true
				}
			}
			if right-left-1 > res {
				res = right - left - 1
			}
		}
	}

	return res
}
