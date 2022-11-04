package mario

var (
	n        = 9
	colMap   = make(map[int]struct{}, n)
	leftMap  = make(map[int]struct{}, n) // y = x + b1
	rightMap = make(map[int]struct{}, n) // y = -x + b2
)

// totalNQueens, 1 <= n <= 9
func totalNQueens(n int) int {
	return dfs(make([]int, n), 0) // slice index: row, value: col of queen
}

func dfs(queens []int, row int) int {
	if row >= len(queens) {
		return 1
	}

	res := 0
	for col := 0; col < len(queens); col++ {
		if _, ok := colMap[col]; ok {
			continue
		}
		if _, ok := leftMap[col-row]; ok {
			continue
		}
		if _, ok := rightMap[col+row]; ok {
			continue
		}

		queens[row] = col
		colMap[col] = struct{}{}
		leftMap[col-row] = struct{}{}
		rightMap[col+row] = struct{}{}

		res += dfs(queens, row+1)

		queens[row] = -1
		delete(colMap, col)
		delete(leftMap, col-row)
		delete(rightMap, col+row)
	}

	return res
}
