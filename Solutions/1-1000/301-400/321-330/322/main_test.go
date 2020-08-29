package mario

import "testing"

type In struct {
	IntSlice0 []int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{[]int{1, 2, 5}, 11}, 3},
	{In{[]int{2}, 3}, -1},
}

func TestCoinChange(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if coinChange(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("coin change test failed on case: %d\n", i)
		}
	}
}
