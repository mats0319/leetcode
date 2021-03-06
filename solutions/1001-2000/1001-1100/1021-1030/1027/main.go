package mario

func uniqueOccurrences(arr []int) bool {
	frequency := make(map[int]int) // elem: appear times
	for _, v := range arr {
		frequency[v]++
	}

	isUnique := true
	elem := make(map[int]int) // appear times: last elem
	for key, value := range frequency {
		if _, ok := elem[value]; ok {
			isUnique = false
			break
		}

		elem[value] = key
	}

	return isUnique
}
