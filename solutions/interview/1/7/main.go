package mario

// matrix.length = matrix[0].length
func rotate(matrix [][]int) {
	n := len(matrix)

	support := make([][]int, n)
	for i := range support {
		support[i] = make([]int, n)
	}

	for i := range matrix {
		for j := range matrix[i] {
			support[j][n-1-i] = matrix[i][j]
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = support[i][j]
		}
	}

	return
}
