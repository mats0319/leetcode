package mario

type Wrapper struct {
	Weight int // distance
	X      int
	Y      int
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	set := make([]*Wrapper, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			set = append(set, &Wrapper{
				Weight: calcDistance(i, j, r0, c0),
				X:      i,
				Y:      j,
			})
		}
	}

	QuickSort(set)

	res := make([][]int, 0, len(set))
	for _, v := range set {
		res = append(res, []int{v.X, v.Y})
	}

	return res
}

func calcDistance(x0, y0, x1, y1 int) int {
	x := x0 - x1
	if x < 0 {
		x = -x
	}

	y := y0 - y1
	if y < 0 {
		y = -y
	}

	return x + y
}

func QuickSort(is []*Wrapper) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i].Weight > is[0].Weight {
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
