package mario

var (
	res          = make([][]string, 0)
	invalidCols  = make(map[int]struct{})
	invalidLeft  = make(map[int]struct{}) // calc: x+y
	invalidRight = make(map[int]struct{}) // calc: x-y
)

func solveNQueens(n int) [][]string {
	res = make([][]string, 0) // for leetcode

	queens := make([]int, n) // keep index of queens on each row
	for i := range queens {
		queens[i] = -1
	}

	dfs(queens, 0)

	return res
}

func dfs(queens []int, row int) {
	if row == len(queens) {
		res = append(res, formatSolution(queens))
		return
	}

	for col := 0; col < len(queens); col++ {
		if _, ok := invalidCols[col]; ok {
			continue
		}
		if _, ok := invalidLeft[row+col]; ok {
			continue
		}
		if _, ok := invalidRight[row-col]; ok {
			continue
		}

		queens[row] = col
		invalidCols[col] = struct{}{}
		invalidLeft[row+col] = struct{}{}
		invalidRight[row-col] = struct{}{}

		dfs(queens, row+1)

		queens[row] = -1
		delete(invalidCols, col)
		delete(invalidLeft, row+col)
		delete(invalidRight, row-col)
	}
}

func formatSolution(queens []int) []string {
	solution := make([]string, 0, len(queens))

	for i := range queens {
		byteSlice := make([]byte, len(queens))
		for j := range byteSlice {
			if j == queens[i] {
				byteSlice[j] = 'Q'
			} else {
				byteSlice[j] = '.'
			}
		}

		solution = append(solution, string(byteSlice))
	}

	return solution
}
