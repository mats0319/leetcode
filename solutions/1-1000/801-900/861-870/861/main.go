package mario

func matrixScore(A [][]int) int {
	for i := range A {
		if A[i][0] != 0 {
			continue
		}

		for j := range A[i] {
			A[i][j] = (A[i][j] + 1) % 2
		}
	}

	sum := 0
	for i, flag := len(A[0])-1, 1; i >= 0; i-- {
		count0 := 0
		count1 := 0
		for _, v := range A {
			if v[i] == 0 {
				count0++
			} else if v[i] == 1 {
				count1++
			}
		}

		sum += max(count0, count1) * flag
		flag *= 2
	}

	return sum
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
