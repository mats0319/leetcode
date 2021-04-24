package mario

type lineWrapper struct {
	k   float64
	num int // num of points
}

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	max := 2
	for i := 0; i < len(points)-1; i++ {
		lines := []lineWrapper{{num: 1}} // lines[0] express k is +max(not exist)
		p1 := []float64{float64(points[i][0]), float64(points[i][1])}

		for j := i + 1; j < len(points); j++ {
			p2 := []float64{float64(points[j][0]), float64(points[j][1])}

			if p1[0] == p2[0] {
				lines[0].num++
				continue
			}

			k := (p2[1] - p1[1]) / (p2[0] - p1[0])
			index := findIndex(lines, k)
			if index < 0 {
				lines = append(lines, lineWrapper{k, 2})
			} else {
				lines[index].num++

				if max < lines[index].num {
					max = lines[index].num
				}
			}
		}
	}

	return max
}

func findIndex(arr []lineWrapper, target float64) int {
	index := -1
	for i := range arr {
		if arr[i].k == target {
			index = i
			break
		}
	}

	return index
}
