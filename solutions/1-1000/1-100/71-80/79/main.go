package mario

func exist(board [][]byte, word string) bool {
	isExist := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				board[i][j] = '.'

				if dfs(board, word[1:], i, j) {
					isExist = true
					break
				}

				board[i][j] = word[0]
			}
		}
	}

	return isExist
}

func dfs(board [][]byte, word string, x, y int) bool {
	if len(word) < 1 {
		return true
	}

	t := board[x][y]

	board[x][y] = '.'

	if x-1 >= 0 &&
		board[x-1][y] == word[0] &&
		dfs(board, word[1:], x-1, y) {
		return true
	}
	if y-1 >= 0 &&
		board[x][y-1] == word[0] &&
		dfs(board, word[1:], x, y-1) {
		return true
	}
	if x+1 < len(board) &&
		board[x+1][y] == word[0] &&
		dfs(board, word[1:], x+1, y) {
		return true
	}
	if y+1 < len(board[x]) &&
		board[x][y+1] == word[0] &&
		dfs(board, word[1:], x, y+1) {
		return true
	}

	board[x][y] = t

	return false
}
