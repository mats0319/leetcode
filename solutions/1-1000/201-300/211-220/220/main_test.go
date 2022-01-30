package mario

import "testing"

type In struct {
	Int0      int
	Int1      int
	IntSlice0 []int
}

var testCase = []struct {
	In     In
	Expect bool
}{
	// test cases here
	{In{0, 3, []int{1, 2, 3, 1}}, true},
	{In{2, 1, []int{1, 0, 1, 1}}, true},
	{In{3, 2, []int{1, 5, 9, 1, 5, 9}}, false},
}

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if containsNearbyAlmostDuplicate(tcs[i].In.IntSlice0, tcs[i].In.Int1, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("contains nearby almost duplicate test failed on case: %d\n", i)
		}
	}
}
