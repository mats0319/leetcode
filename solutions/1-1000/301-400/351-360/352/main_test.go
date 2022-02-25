package mario

import "testing"

func TestSummaryRanges_Case1(t *testing.T) {
	summaryRangesIns := Constructor()

	res := make([][][]int, 0)

	summaryRangesIns.AddNum(6)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(6)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(0)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(4)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(8)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(7)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(6)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(4)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(7)
	res = append(res, summaryRangesIns.GetIntervals())
	summaryRangesIns.AddNum(5)
	res = append(res, summaryRangesIns.GetIntervals())

	expected := [][][]int{
		{{6, 6}},
		{{6, 6}},
		{{0, 0}, {6, 6}},
		{{0, 0}, {4, 4}, {6, 6}},
		{{0, 0}, {4, 4}, {6, 6}, {8, 8}},
		{{0, 0}, {4, 4}, {6, 8}},
		{{0, 0}, {4, 4}, {6, 8}},
		{{0, 0}, {4, 4}, {6, 8}},
		{{0, 0}, {4, 4}, {6, 8}},
		{{0, 0}, {4, 8}},
	}

	if len(res) != len(expected) {
		t.Errorf("unmatched res length, get: %d, want: %d\n", len(res), len(expected))
	}

	for i := range res {
		if !compareOnTwoDimensionalIntSlice(res[i], expected[i]) {
			t.Errorf("unexpected res, index: %d, get: %v, want: %v\n", i, res[i], expected[i])
		}
	}
}

// compareOnTwoDimensionalIntSlice return if a and b is strictly equal
func compareOnTwoDimensionalIntSlice(a, b [][]int) bool {
    if len(a) != len(b) {
        return false
    }

    isEqual := true
    for i := range a {
        if !compareOnIntSlice(a[i], b[i]) {
            isEqual = false
            break
        }
    }

    return isEqual
}

func compareOnIntSlice(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    isEqual := true
    for i := range a {
        if a[i] != b[i] {
            isEqual = false
            break
        }
    }

    return isEqual
}
