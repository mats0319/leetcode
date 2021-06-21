package mario

func setZeroes(matrix [][]int) {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return
	}

	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < len(row); i++ {
		if row[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < len(col); j++ {
		if col[j] {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}

	return
}
