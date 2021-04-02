package mario

// matrix is always not nil
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])

	res := make([]int, 0, row * col)

	for i := 0; i < (min(row, col)+1)/2; i++ {
		// top: (i, i) -> (i, col-1-i) +
		for c := i; c <= col-1-i; c++ {
			res = append(res, matrix[i][c])
		}

		// right: (i+1, col-1-i) -> (row-1-i, col-1-i) +
		for r := i+1; r <= row-1-i; r++ {
			res = append(res, matrix[r][col-1-i])
		}

		if len(res) == cap(res) {
			break
		}

		// bottom: (row-1-i, col-1-i-1) -> (row-1-i, i) -
		for c := col-2-i; c >= i; c-- {
			res = append(res, matrix[row-1-i][c])
		}

		// left: (row-1-i-1, i) -> (i+1, i) -
		for r := row-2-i; r >= i+1; r-- {
			res = append(res, matrix[r][i])
		}
	}

	return res
}

func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
