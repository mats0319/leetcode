package mario

func maximalRectangle(matrix [][]byte) (result int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// matrix[i][j].height
	height := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := range height {
			height[j] = (height[j] + 1) * int(matrix[i][j]-'0') // update height

			if height[j] > 0 {
				minHeight := height[j]
				for k := j; k >= 0 && height[k] > 0; k-- { // loop valid width
					minHeight = min(minHeight, height[k])
					s := minHeight * (j - k + 1)

					if result < s {
						result = s
					}
				}
			}
		}
	}

	return
}

func min(a, b int) (result int) {
	if a < b {
		result = a
	} else {
		result = b
	}

	return
}
