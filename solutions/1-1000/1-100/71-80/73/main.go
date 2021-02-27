package mario

func setZeroes(matrix [][]int) {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return
	}

	row := make([]bool, len(matrix))
	column := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				column[j] = true
			}
		}
	}

	for i := 0; i < len(row); i++ {
		if !row[i] {
			continue
		}

		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = 0
		}
	}

	for i := 0; i < len(column); i++ {
		if !column[i] {
			continue
		}

		for j := 0; j < len(matrix); j++ {
			matrix[j][i] = 0
		}
	}

	return
}
