package mario

type Appear struct {
	AppearTime int
	FirstIndex int
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]*Appear)
	maxAppearTime := 0
	minLength := len(nums)
	for i, value := range nums {
		t, ok := m[value]
		if !ok {
			t = &Appear{FirstIndex: i}
		}
		t.AppearTime++

		m[value] = t

		if maxAppearTime < t.AppearTime {
			maxAppearTime = t.AppearTime
			minLength = i - t.FirstIndex + 1
		} else if maxAppearTime == t.AppearTime && minLength > i-t.FirstIndex+1 {
			minLength = i - t.FirstIndex + 1
		}
	}

	return minLength
}
