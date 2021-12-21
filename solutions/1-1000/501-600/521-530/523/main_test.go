package mario

import "testing"

type In struct {
	IntSlice0 []int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect bool
}{
	// test cases here
	{In{[]int{5, 0, 0, 0}, 3}, true},
	{In{[]int{23, 2, 4, 6, 6}, 7}, true},
	{In{[]int{23, 2, 4, 6, 7}, 6}, true},
}

func TestCheckSubarraySum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if checkSubarraySum(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("check subarray sum test failed on case: %d\n", i)
		}
	}
}
