package mario

// matrix always not nil
func searchMatrix(matrix [][]int, target int) bool {
	isExist := false
	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[0]); {
		if matrix[i][j] == target {
			isExist = true
			break
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}

	return isExist
}
