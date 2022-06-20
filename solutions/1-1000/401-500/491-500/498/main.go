package mario

// findDiagonalOrder mat each dimension is not nil
func findDiagonalOrder(mat [][]int) []int {
	x := len(mat)
	y := len(mat[0])

	res := make([]int, 0, x*y)
	for i := 0; i < x+y-1; i++ {
		res = append(res, traversalMatrix(mat, i, i%2 == 1)...)
	}

	return res
}

func traversalMatrix(mat [][]int, sum int, isPosDirection bool) []int {
	x := small(sum, len(mat)-1)
	y := sum - x

	res := make([]int, 0)
	for x >= 0 && y < len(mat[0]) {
		res = append(res, mat[x][y])

		x--
		y++
	}

	if !isPosDirection {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}

	return res
}

func small(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
