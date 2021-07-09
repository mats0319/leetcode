package mario

type line struct {
	x int
	y int
}

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	res := 0
ALL:
	for i := 0; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]

		for j := i + 1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]

			k1 := calcK(x1, y1, x2, y2)
			count := 2

			for k := j + 1; k < len(points); k++ {
				x3, y3 := points[k][0], points[k][1]

				k2 := calcK(x1, y1, x3, y3)
				if k1 == k2 {
					count++
				}
			}

			if count > res {
				res = count
			}

			if res >= len(points)-i-1 {
				break ALL
			}
		}
	}

	return res
}

// calcK make sure x is positive
func calcK(x1, y1, x2, y2 int) line {
	if x1 == x2 {
		return line{x: 0, y: x1}
	} else if y1 == y2 {
		return line{x: y1, y: 0}
	}

	x := x1 - x2
	y := y1 - y2
	if x < 0 {
		x *= -1
		y *= -1
	}

	g := gcd(pos(x), pos(y))

	return line{x: x / g, y: y / g}
}

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a < b {
		a, b = b, a
	}

	for {
		mod := a % b
		if mod == 0 {
			break
		}

		a = b
		b = mod
	}

	return b
}

func pos(n int) int {
	if n < 0 {
		n *= -1
	}

	return n
}
