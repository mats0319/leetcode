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
	{In{[]int{6, 3, 5, 7, 2, 3, 3, 8, 2, 4}, 2}, 5},
}

func TestFindPairs(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findPairs(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("find pairs test failed on case: %d\n", i)
		}
	}
}
