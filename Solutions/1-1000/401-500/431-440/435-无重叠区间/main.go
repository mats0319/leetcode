package mario

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var (
		end   = intervals[0][1]
		count = 1 // count counts no overlap intervals' number
	)
	for i := 1; i < len(intervals); i++ {
		if end <= intervals[i][0] {
			end = intervals[i][1]
			count++
		}
	}

	return len(intervals) - count
}

// modify according to quick sort.
func sortByEnd(is [][]int) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i][1] > is[0][1] {
				big++
			} else {
				small++
				if big != 0 {
					is[i], is[small] = is[small], is[i]
				}
			}
		}
		if small != 0 {
			is[0], is[small] = is[small], is[0]
		}
	}

	sortByEnd(is[:small])
	sortByEnd(is[small+1:])

	return
}
