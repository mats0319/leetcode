package mario

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)

	mergeStart := 0
	i := 0
	for i < len(intervals) && newInterval[0] > intervals[i][1] {
		i++
	}
	res = append(res, intervals[:i]...)
	if i >= len(intervals) { // intervals.length = 0, exit here
		res = append(res, newInterval)
		return res
	}

	mergeStart = small(intervals[i][0], newInterval[0])
	for i < len(intervals) && newInterval[1] >= intervals[i][0] {
		i++
	}

	mergeEnd := newInterval[1]
	if i-1 >= 0 && newInterval[1] < intervals[i-1][1] {
		mergeEnd = intervals[i-1][1]
	}

	res = append(res, []int{mergeStart, mergeEnd})

	res = append(res, intervals[i:]...)

	return res
}

func small(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
