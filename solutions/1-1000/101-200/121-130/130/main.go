package mario

type point struct {
	x int
	y int
}

const marginal = '1'

// board is always not nil
func solve(board [][]byte) {
	marginalPoints := traversalMarginal(board)
	for len(marginalPoints) > 0 {
		p := marginalPoints[0]
		marginalPoints = marginalPoints[1:]

		board[p.x][p.y] = marginal
		// todo: four direction check
	}

	// todo: traversal and change 'O' to 'X', '' to 'O'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			switch board[i][j] {
			case 'O':
			}
		}
	}

	return
}

func traversalMarginal(board [][]byte) []*point {
	row := len(board)
	col := len(board[0])

	res := make([]*point, 0, (row+col)*2)
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			res = append(res, &point{i, 0})
		}
		if board[i][col-1] == 'O' {
			res = append(res, &point{i, col - 1})
		}
	}

	for j := 1; j < col-1; j++ { // 防止角落重复
		if board[0][j] == 'O' {
			res = append(res, &point{0, j})
		}
		if board[row-1][j] == 'O' {
			res = append(res, &point{row - 1, j})
		}
	}

	return res
}
