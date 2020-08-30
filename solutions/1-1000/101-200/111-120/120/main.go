package mario

func minimumTotal(triangle [][]int) (result int) {
	for i := 1; i < len(triangle); i++ {
		triangle[i][i] += triangle[i-1][i-1]
		for j := i - 1; j > 0; j-- {
			triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
		}
		triangle[i][0] += triangle[i-1][0]
	}

	result = triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle); i++ {
		if result > triangle[len(triangle)-1][i] {
			result = triangle[len(triangle)-1][i]
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
