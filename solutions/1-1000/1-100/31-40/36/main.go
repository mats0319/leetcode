package mario

func isValidSudoku(board [][]byte) bool {
	isValid := true

RC:
	for i := 0; i < len(board); i++ {
		existRow := make([]bool, len(board)+1)
		existCol := make([]bool, len(board[i])+1)

		for j := 0; j < len(board[i]); j++ {
			if '0' <= board[i][j] && board[i][j] <= '9' {
				if existRow[board[i][j]-'0'] {
					isValid = false
					break RC
				}

				existRow[board[i][j]-'0'] = true
			}

			if '0' <= board[j][i] && board[j][j] <= '9' {
				if existCol[board[j][i]-'0'] {
					isValid = false
					break RC
				}

				existCol[board[j][i]-'0'] = true
			}
		}
	}

	if !isValid {
		return false
	}

GRID:
	for rowCoefficient := 0; rowCoefficient < len(board)/3; rowCoefficient++ {
		for colCoefficient := 0; colCoefficient < len(board[rowCoefficient])/3; colCoefficient++ {
			existGrid := make([]bool, len(board)+1)

			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					index := int(board[row+rowCoefficient*3][col+colCoefficient*3] - '0')
					if 0 <= index && index <= 9 {
						if existGrid[index] {
							isValid = false
							break GRID
						}

						existGrid[index] = true
					}
				}
			}
		}
	}

	return isValid
}
