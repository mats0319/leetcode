package mario

func findDiagonalOrder(mat [][]int) []int {
	if len(mat) < 1 || len(mat[0]) < 1 {
		return nil
	}

	res := make([]int, 0, len(mat)*len(mat[0]))
	for sum := 0; sum < len(mat)+len(mat[0])-1; sum++ {
		if sum%2 == 0 { // x max -> 0
			for x, y := sum, 0; x >= 0 && y < len(mat[0]); x, y = x-1, y+1 {
				if x >= len(mat) {
					continue
				}
				res = append(res, mat[x][y])
			}
		} else {
			for x, y := 0, sum; x < len(mat) && y >= 0; x, y = x+1, y-1 {
				if y >= len(mat[0]) {
					continue
				}
				res = append(res, mat[x][y])
			}
		}
	}

	return res
}
