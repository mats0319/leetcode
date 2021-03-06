package mario

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) < 2 {
		return len(envelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			(envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	length := make([]int, len(envelopes))
	max := 0
	for i := 0; i < len(envelopes); i++ {
		maxLength := 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] && maxLength < length[j]+1 {
				maxLength = length[j] + 1
			}
		}

		length[i] = maxLength

		if max < length[i] {
			max = length[i]
		}
	}

	return max
}
