package mario

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	var result byte = '0'
Find1:
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				result = '1'
				break Find1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				matrix[i][j] = min(min(matrix[i-1][j], matrix[i][j-1]), matrix[i-1][j-1]) + 1
				if result < matrix[i][j] {
					result = matrix[i][j]
				}
			}
		}
	}

	return int(result-'0') * int(result-'0')
}

func min(a, b byte) (result byte) {
	if a < b {
		result = a
	} else {
		result = b
	}

	return
}
