package mario

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

// board always not nil
func gameOfLife(board [][]int) {
	// alive, around alive != 2 or 3, die
	// die, around alive = 3, alive

	// base case: 0: die, 1: alive;
	// further case: 2: die from alive, -1: alive from die.
	// rule: >= 1: alive, <= 0: die
	// next traversal: 2 -> 0, -1 -> 1

	row := len(board)
	col := len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			aliveCount := 0
			for _, v := range directions {
				x := i + v[0]
				y := j + v[1]
				if 0 <= x && x < row &&
					0 <= y && y < col &&
					board[x][y] >= 1 {
					aliveCount++
				}
			}

			if board[i][j] == 0 && aliveCount == 3 {
				board[i][j] = -1
			}
			if board[i][j] == 1 && aliveCount != 2 && aliveCount != 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}

	return
}
