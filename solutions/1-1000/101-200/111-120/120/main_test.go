package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
}

func TestMinimumTotal(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minimumTotal(tcs[i].In) != tcs[i].Expect {
			t.Errorf("minimum total test failed on case: %d\n", i)
		}
	}
}
