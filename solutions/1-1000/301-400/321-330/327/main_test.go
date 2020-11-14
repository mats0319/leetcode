package mario

import "testing"

type In struct {
	Int1      int
	Int0      int
	IntSlice0 []int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{-2, 2, []int{-2, 5, -1}},  3},
}

func TestCountRangeSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if countRangeSum(tcs[i].In.IntSlice0, tcs[i].In.Int1, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("count range sum test failed on case: %d\n", i)
		}
	}
}
