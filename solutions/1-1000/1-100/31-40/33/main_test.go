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
	{In{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
	{In{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
	{In{[]int{1}, 0}, -1},
	{In{[]int{3, 1}, 0}, -1},
	{In{[]int{5, 1, 3}, 2}, -1},
}

func TestSearch(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if search(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("search test failed on case: %d\n", i)
		}
	}
}
