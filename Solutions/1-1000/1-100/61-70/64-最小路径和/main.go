package mario

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) (result int) {
	if a < b {
		result = a
	} else {
		result = b
	}

	return
}
