package mario

func merge(A []int, m int, B []int, n int) {
	indexA := m - 1
	indexB := n - 1

	for i := m + n - 1; i >= 0 && indexA != -1 && indexB != -1; i-- {
		if A[indexA] > B[indexB] {
			A[i] = A[indexA]
			indexA--
		} else {
			A[i] = B[indexB]
			indexB--
		}
	}

	if indexA == -1 && indexB != -1 {
		for ; indexB >= 0; indexB-- {
			A[indexB] = B[indexB]
		}
	}

	return
}
