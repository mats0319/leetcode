package mario

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		j := i + 1
		for ; j < len(intervals) && intervals[j][0] <= end; j++ {
			if end < intervals[j][1] {
				end = intervals[j][1]
			}
		}

		res = append(res, []int{start, end})

		i = j - 1
	}

	return res
}
