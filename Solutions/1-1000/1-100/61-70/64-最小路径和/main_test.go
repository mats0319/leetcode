package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
}

func TestMinPathSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minPathSum(tcs[i].In) != tcs[i].Expect {
			t.Errorf("min path sum test failed on case: %d", i)
		}
	}
}
