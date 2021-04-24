package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	// test cases here
	{[][]int{
		{100, -42, -46, -41},
		{31, 97, 10, -10},
		{-58, -51, 82, 89},
		{51, 81, 69, -51},
	}, -36},
}

func TestMinFallingPathSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minFallingPathSum(tcs[i].In) != tcs[i].Expect {
			t.Errorf("min falling path sum test failed on case: %d\n", i)
		}
	}
}
