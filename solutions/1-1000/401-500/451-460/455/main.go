package mario

func findContentChildren(g []int, s []int) int {
	QuickSort(g)
	QuickSort(s)

	res := 0
	index := 0
	for i := 0; i < len(g) && index < len(s); i++ {
		for index < len(s) {
			if s[index] >= g[i] {
				index++
				res++
				break
			}

			index++
		}
	}

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
