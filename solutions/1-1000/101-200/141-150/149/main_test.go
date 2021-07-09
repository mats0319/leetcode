package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	// test cases here
	{[][]int{
		{0, 0},
		{4, 5},
		{7, 8},
		{8, 9},
		{5, 6},
		{3, 4},
		{1, 1},
	}, 5},
	{[][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}, 4},
	{[][]int{
		{2, 3},
		{3, 3},
		{-5, 3},
	}, 3},
}

func TestMaxPoints(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxPoints(tcs[i].In) != tcs[i].Expect {
			t.Errorf("max points test failed on case: %d\n", i)
		}
	}
}
