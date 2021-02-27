package mario

func isToeplitzMatrix(matrix [][]int) bool {
	isValid := true
ALL:
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				isValid = false
				break ALL
			}
		}
	}

	return isValid
}
