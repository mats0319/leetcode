package mario

func numEquivDominoPairs(dominoes [][]int) int {
	times := make([]int, 100)
	count := 0
	for _, pair := range dominoes {
		count += times[pair[0]*10+pair[1]]
		if pair[0] != pair[1] {
            count += times[pair[1]*10+pair[0]]
        }

		times[pair[0]*10+pair[1]]++
	}

	return count
}
