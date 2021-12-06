package mario

func distributeCandies(candyType []int) int {
	typeCountMap := make(map[int]int) // type - count
	for i := range candyType {
		typeCountMap[candyType[i]]++
	}

	return min(len(typeCountMap), len(candyType)/2)
}

func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return res
}
