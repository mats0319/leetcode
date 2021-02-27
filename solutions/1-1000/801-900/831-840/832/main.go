package mario

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		left := 0
		right := len(A[i]) - 1
		for left <= right {
			A[i][left], A[i][right] = 1-A[i][right], 1-A[i][left]

			left++
			right--
		}
	}

	return A
}
