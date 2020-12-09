package mario

func leastInterval(tasks []byte, n int) int {
	count := [26]int{}
	maxCount := 0
	for i := 0; i < len(tasks); i++ {
		count[int(tasks[i]-'A')]++
		if count[int(tasks[i]-'A')] > maxCount {
			maxCount = count[int(tasks[i]-'A')]
		}
	}

	times := make(map[int]int, 26)
	for _, v := range count {
		times[v]++
	}
	times[0] = 0

	result := (maxCount-1)*(n+1) + times[maxCount]

	// if least interval is not valid(some chars need append in each sub-interval),
	// just re-order origin array is ok
	if len(tasks) > result {
		result = len(tasks)
	}

	return result
}
