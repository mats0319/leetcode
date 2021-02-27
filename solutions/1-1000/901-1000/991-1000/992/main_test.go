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
	{In{[]int{1, 2, 1, 2, 3}, 2}, 7},
	{In{[]int{1, 2, 1, 3, 4}, 3}, 3},
}

func TestSubarraysWithKDistinct(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if subarraysWithKDistinct(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("subarrays with k distinct test failed on case: %d\n", i)
		}
	}
}
