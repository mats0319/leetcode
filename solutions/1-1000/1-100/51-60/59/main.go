package mario

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	count := 1
	for i := 0; i < (n+1)/2; i++ {
		// top: (i, i) -> (i, col-1-i) +
		for c := i; c <= n-1-i; c++ {
			res[i][c] = count
			count++
		}

		// right: (i+1, col-1-i) -> (row-1-i, col-1-i) +
		for r := i + 1; r <= n-1-i; r++ {
			res[r][n-1-i] = count
			count++
		}

		if count > n*n {
			break
		}

		// bottom: (row-1-i, col-1-i-1) -> (row-1-i, i) -
		for c := n - 2 - i; c >= i; c-- {
			res[n-1-i][c] = count
			count++
		}

		// left: (row-1-i-1, i) -> (i+1, i) -
		for r := n - 2 - i; r >= i+1; r-- {
			res[r][i] = count
			count++
		}
	}

	return res
}
