package mario

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))

	for i := range list1 {
		m[list1[i]] = i
	}

	res := make([]string, 0, 1)
	minSum := len(list1) + len(list2)
	for i := range list2 {
		index, ok := m[list2[i]]
		if !ok {
			continue
		}

		if minSum > index+i {
			res = []string{list2[i]}
			minSum = index + i
		} else if minSum == index+i {
			res = append(res, list2[i])
		}
	}

	return res
}
