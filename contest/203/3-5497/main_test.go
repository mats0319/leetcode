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
	{In{[]int{3,5,1,2,4}, 1}, 4},
	{In{[]int{3,1,5,4,2}, 2}, -1},
	{In{[]int{1}, 1}, 1},
	{In{[]int{2,1}, 2}, 2},
	{In{[]int{1,4,2,3}, 4}, 4},
}

func TestFindLatestStep(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findLatestStep(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("find latest step test failed on case: %d\n", i)
		}
	}
}
