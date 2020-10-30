package mario

func islandPerimeter(grid [][]int) int {
	sum := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				sum += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					sum--
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					sum--
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					sum--
				}
				if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
					sum--
				}
			}
		}
	}

	return sum
}
