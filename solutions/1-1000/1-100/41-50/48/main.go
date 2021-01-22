package mario

func rotate(matrix [][]int) {
	n := len(matrix)

	support := make([][]int, n)
	for i := range support {
		support[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			support[j][n-i-1] = matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = support[i][j]
		}
	}

	return
}
