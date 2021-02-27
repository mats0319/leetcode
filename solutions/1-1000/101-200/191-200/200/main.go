package mario

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}

			count++
			dfs(grid, i, j)
		}
	}

	return count
}

func dfs(grid [][]byte, x, y int) {
	grid[x][y] = '0'

	if x-1 >= 0 && grid[x-1][y] == '1' {
		dfs(grid, x-1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == '1' {
		dfs(grid, x, y-1)
	}
	if x+1 < len(grid) && grid[x+1][y] == '1' {
		dfs(grid, x+1, y)
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == '1' {
		dfs(grid, x, y+1)
	}

	return
}
