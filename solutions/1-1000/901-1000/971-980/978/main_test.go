package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}, 5},
	{[]int{1, 2, 3, 4}, 2},
	{[]int{1}, 1},
	{[]int{9, 9}, 1},
	{[]int{2, 0, 2, 4, 2, 5, 0, 1, 2, 3}, 6},
}

func TestMaxTurbulenceSize(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxTurbulenceSize(tcs[i].In) != tcs[i].Expect {
			t.Errorf("max turbulence size test failed on case: %d\n", i)
		}
	}
}
