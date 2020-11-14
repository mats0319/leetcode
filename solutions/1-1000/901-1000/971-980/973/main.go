package mario

type SquareWithIndex struct {
	Square int
	Index  int
}

func kClosest(points [][]int, K int) [][]int {
	if len(points) == K {
		return points
	}

	square := make([]*SquareWithIndex, 0, len(points))
	for i := range points {
		p := points[i]
		square = append(square, &SquareWithIndex{
			Square: p[0]*p[0] + p[1]*p[1],
			Index:  i,
		})
	}

	QuickSort(square)

	res := make([][]int, 0, K)
	for i := 0; i < K; i++ {
		res = append(res, points[square[i].Index])
	}

	return res
}

func QuickSort(is []*SquareWithIndex) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i].Square > is[0].Square {
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
