package mario

func solveSudoku(board [][]byte) {
	_ = dfs(board, 0, 0)
}

func dfs(board [][]byte, i int, j int) bool {
	if i > 8 {
		return true
	}

	ni, nj := nextItem(i, j)

	if board[i][j] != '.' { // not empty
		return dfs(board, ni, nj)
	}

	options := getOptions(board, i, j)
	res := false
	for k := range options {
		board[i][j] = options[k]

		res = dfs(board, ni, nj)
		if res {
			break
		}

		board[i][j] = '.'
	}

	return res
}

func getOptions(board [][]byte, i int, j int) []byte {
	invalidOptions := make([]bool, 9) // define on length

	for k := 0; k < 9; k++ {
		if board[i][k] != '.' {
			invalidOptions[board[i][k]-'1'] = true // row i
		}
		if board[k][j] != '.' {
			invalidOptions[board[k][j]-'1'] = true // col j
		}
	}

	row := i / 3
	col := j / 3

	for r := row * 3; r < row*3+3; r++ {
		for c := col * 3; c < col*3+3; c++ {
			if board[r][c] != '.' {
				invalidOptions[board[r][c]-'1'] = true // grid
			}
		}
	}

	res := make([]byte, 0, 9)
	for k := range invalidOptions {
		if !invalidOptions[k] {
			res = append(res, byte(k)+'1')
		}
	}

	return res
}

func nextItem(i, j int) (int, int) {
	if j == 8 {
		i++
		j = 0
	} else {
		j++
	}

	return i, j
}
