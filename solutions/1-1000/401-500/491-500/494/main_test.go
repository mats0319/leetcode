package mario

import "testing"

type In struct {
	Int0      int
	IntSlice0 []int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{3, []int{1, 1, 1, 1, 1}}, 5},
	{In{1, []int{1}}, 1},
}

func TestFindTargetSumWays(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findTargetSumWays(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("find target sum ways test failed on case: %d\n", i)
		}
	}
}
