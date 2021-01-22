package mario

func computeArea(
	A int, B int, C int, D int,
	E int, F int, G int, H int,
) int {
	area := (C-A)*(D-B) + (G-E)*(H-F)
	covered := calcOverlapping(A, C, E, G) * calcOverlapping(B, D, F, H)

	return area - covered
}

func calcOverlapping(startA, endA, startB, endB int) int {
	if startB > endA || startA > endB {
		return 0
	}

	return min(endA, endB) - max(startA, startB)
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}

func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
