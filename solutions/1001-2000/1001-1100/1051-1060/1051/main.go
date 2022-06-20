package mario

import "sort"

func heightChecker(heights []int) int {
	expected := make([]int, 0, len(heights))
	expected = append(expected, heights...)

	sort.Ints(expected)

	count := 0
	for i := range expected {
		if expected[i] != heights[i] {
			count++
		}
	}

	return count
}
