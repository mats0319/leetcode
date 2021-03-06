package mario

func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	QuickSort(points)

	count := 0
	lastInterval := points[0]
	for i := 1; i < len(points); i++ {
		currInterval := points[i]
		if currInterval[0] == lastInterval[0] && currInterval[1] < lastInterval[1] {
			lastInterval[1] = currInterval[1]
		} else if currInterval[0] != lastInterval[0] && currInterval[0] < lastInterval[1] {
			lastInterval = []int{currInterval[0], small(currInterval[1], lastInterval[1])}
		} else if currInterval[0] > lastInterval[1] {
			count++
			lastInterval = currInterval
		}
	}

	return count + 1
}

func QuickSort(is [][]int) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i][0] > is[0][0] {
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

func small(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
