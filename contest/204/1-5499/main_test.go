package mario

import "testing"

type In struct {
	IntSlice0 []int
	Int0      int
	Int1      int
}

var testCase = []struct {
	In     In
	Expect bool
}{
	// test cases here
	{In{[]int{1, 2, 4, 4, 4, 4}, 1, 3}, true},
	{In{[]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2}, true},
	{In{[]int{1, 2, 1, 2, 1, 3}, 2, 3}, false},
	{In{[]int{1, 2, 3, 1, 2}, 2, 2}, false},
	{In{[]int{2, 2, 2, 2}, 2, 3}, false},
}

func TestContainsPattern(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if containsPattern(tcs[i].In.IntSlice0, tcs[i].In.Int1, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("contains pattern test failed on case: %d\n", i)
		}
	}
}
