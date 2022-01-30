package mario

import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)

	l, r := 0, len(piles)-1
	res := 0
	for l <= r {
		r--
		res += piles[r]

		l++
		r--
	}

	return res
}
