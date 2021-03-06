package mario

import "testing"

type In struct {
	IntSlice0 [][]int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect bool
}{
	// test cases here
	{In{[][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5}, true},
}

func TestSearchMatrix(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if searchMatrix(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("search matrix test failed on case: %d\n", i)
		}
	}
}
