package mario

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	res := make([][]int, 0, numRows)
	res = append(res, []int{1})

	for i := 2; i <= numRows; i++ {
		line := make([]int, i)
		line[0] = 1
		line[i-1] = 1

		for j := 1; j < i-1; j++ {
			line[j] = res[i-2][j-1] + res[i-2][j]
		}

		res = append(res, line)
	}

	return res
}
