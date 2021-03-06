package mario

func longestOnes(A []int, K int) int {
	maxLength := 0
	{
		zeroCount := 0
		left, right := 0, 0
		for ; right < len(A); right++ {
			if A[right] == 1 {
				continue
			}

			zeroCount++
			if zeroCount <= K {
				continue
			}

			if maxLength < right-left { // right - left + 1 - 1, decrease current A[right]
				maxLength = right - left
			}

			for A[left] == 1 {
				left++
			}

			left++
			zeroCount--
		}

		if maxLength < right-left {
			maxLength = right - left
		}
	}

	return maxLength
}
