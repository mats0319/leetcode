package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
	{[][]int{{1}}, 0},
	{[][]int{{1, 0}}, 0},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if uniquePathsWithObstacles(tcs[i].In) != tcs[i].Expect {
			t.Errorf("unique paths with obstacles test failed on case: %d", i)
		}
	}
}
