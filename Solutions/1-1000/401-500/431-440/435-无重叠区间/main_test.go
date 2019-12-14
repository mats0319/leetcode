package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
	{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
	{[][]int{{1, 2}, {2, 3}}, 0},
}

func TestEraseOverlapIntervals(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if eraseOverlapIntervals(tcs[i].In) != tcs[i].Expect {
			t.Errorf("erase overlap intervals test failed on case: %d", i)
		}
	}
}
