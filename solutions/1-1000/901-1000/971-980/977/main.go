package mario

func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	for i, e := range A {
		res[i] = e * e
	}

	QuickSort(res)

	return res
}

// QuickSort quick sort
func QuickSort(is []int) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i] > is[0] {
				big++
			} else {
				small++
				if big != 0 {
					is[i], is[small] = is[small], is[i]
				}
			}
		}
		if small != 0 {
			is[0], is[small] = is[small], is[0]
		}
	}

	QuickSort(is[:small])
	QuickSort(is[small+1:])

	return
}
