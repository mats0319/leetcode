package mario

import (
    "fmt"
)

var options [][][]bool

func solveSudoku(board [][]byte) {
    options = make([][][]bool, 9) // for lc
    for i := range options {
        options[i] = make([][]bool, 9)
        for j := range options[i] {
            if board[i][j] == '.' {
                options[i][j] = make([]bool, 9) // means 1-9, if someone appear, set index: 'value-1' = true
            }
        }
    }

    isFinished := false
    for count := 0; !isFinished; count++ {
        if count % 100 == 0 {
            fmt.Println(fmt.Sprintf("> it loops %d times.", count))
        }

        isFinished = true
        for i := range board {
            for j := range board {
                if board[i][j] != '.' {
                    reduceOptions(board, i, j)
                    isFinished = false
                }
            }
        }

        board = writeSolution(board)
    }

    return
}

func reduceOptions(board [][]byte, i, j int) {
    for c, r := range board[i] {
        if board[i][c] == '.' {
            options[i][c][int(r-'1')] = true // use "-'1'" instead of "-'0'-1"
        }
    }
    for c := range board {
        if board[c][j] == '.' {
            options[c][j][int(board[c][j]-'1')] = true
        }
    }
    in, jn := i / 3, j / 3
    for m := in * 3; m < in*3+3; m++ {
        for n := jn * 3; n < jn*3+3; n++ {
            if board[m][n] == '.' {
                options[m][n][int(board[m][n]-'1')] = true
            }
        }
    }

    return
}

func writeSolution(board [][]byte) [][]byte {
    for i := range board {
        for j := range board[i] {
            index := onlyOneOption(i, j)
            if board[i][j] == '.' && index != -1 {
                board[i][j] = byte(index+'0')
            }
        }
    }

    // if one row / column / grid only one position can fill a word, just fill it
    for i := range board {
        count := make([][]int, 9) // count[a]: []int, means in row i, where byte index: 'a+1' may appear
        for j := range board[i] {
            if board[i][j] == '.' {
                for c, v := range options[i][j] {
                    if !v {
                        count[c] = append(count[c], j)
                    }
                }
            }
        }

        for i, pos := range count {
            if len(pos) == 1 { // in row i, if a byte can only appear at somewhere
                board[i][pos[0]] = byte(i+'1')
            }
        }
    }

    for j := range board {
        count := make([][]int, 9)
        for i := range board[j] {
            if board[i][j] == '.' {
                for c, v := range options[i][j] {
                    if !v {
                        count[c] = append(count[c], i)
                    }
                }
            }
        }

        for i, pos := range count {
            if len(pos) == 1 {
                board[pos[0]][j] = byte(i+'1')
            }
        }
    }

    for in := 0; in < 3; in++ {
        for jn := 0; jn < 3; jn++ {
            type point struct {
                X int
                Y int
            }

            count := make([][]*point, 9)
            for i := in*3; i < in*3+3; i++ {
                for j := jn*3; j < jn*3+3; j++ {
                    if board[i][j] == '.' {
                        for c, v := range options[i][j] {
                            if !v {
                                count[c] = append(count[c], &point{i, j})
                            }
                        }
                    }
                }
            }

            for b, pos := range count {
                if len(pos) == 1 {
                    board[pos[0].X][pos[0].Y] = byte(b+'1')
                }
            }
        }
    }

    return board
}

// make sure options[i][j] has initialized
func onlyOneOption(i, j int) int {
    count := 0
    index := -1
    for i, v := range options[i][j] {
        if !v {
            count++
            index = i
            if count > 1 {
                break
            }
        }
    }

    if count != 1 {
        index = -1
    }

    return index
}
