package mario

func numberOfBoomerangs(points [][]int) int {
    res := 0
	for i := range points {
		x1, y1 := points[i][0], points[i][1]
		m := make(map[int]int, len(points)) // distance - points' count
		for j := range points {
			m[calcDistance(x1, y1, points[j][0], points[j][1])]++
		}

		for _, v := range m {
		    if v > 1 {
		        res += v * (v-1)
            }
        }
	}

	return res
}

func calcDistance(x1, y1, x2, y2 int) int {
	x := x1 - x2
	if x < 0 {
		x *= -1
	}

	y := y1 - y2
	if y < 0 {
		y *= -1
	}

	return x*x + y*y
}
