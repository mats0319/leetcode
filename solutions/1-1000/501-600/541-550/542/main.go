package mario

// at least a '0', mat.length / mat[i].length >= 1
func updateMatrix(mat [][]int) [][]int {
	allZero := true
	queue := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
				allZero = false
			}
		}
	}

	if allZero {
		return mat
	}

	for len(queue) > 0 {
		nextQueue := make([][]int, 0)

		for len(queue) > 0 {
			point := queue[0]
			queue = queue[1:]

			x, y := point[0], point[1]
			value := mat[x][y]
			// x-1, y
			if x-1 >= 0 && mat[x-1][y] < 0 {
				nextQueue = append(nextQueue, []int{x - 1, y})
				mat[x-1][y] = value + 1
			}
			// x+1, y
			if x+1 < len(mat) && mat[x+1][y] < 0 {
				nextQueue = append(nextQueue, []int{x + 1, y})
				mat[x+1][y] = value + 1
			}
			// x, y-1
			if y-1 >= 0 && mat[x][y-1] < 0 {
				nextQueue = append(nextQueue, []int{x, y - 1})
				mat[x][y-1] = value + 1
			}
			// x, y+1
			if y+1 < len(mat[0]) && mat[x][y+1] < 0 {
				nextQueue = append(nextQueue, []int{x, y + 1})
				mat[x][y+1] = value + 1
			}
		}

		queue = nextQueue
	}

	return mat
}
