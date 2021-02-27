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
	{In{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2}, 6},
	{In{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3}, 10},
	{In{[]int{0, 0, 0, 1}, 4}, 4},
}

func TestLongestOnes(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestOnes(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("longest ones test failed on case: %d\n", i)
		}
	}
}
