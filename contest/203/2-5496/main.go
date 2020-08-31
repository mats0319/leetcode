package mario

import "github.com/mats9693/utils/sort"

func maxCoins(piles []int) int {
	sort.QuickSort(piles)

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
