package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{1,-2,-3,4}, 4},
	{[]int{0,1,-2,-3,-4}, 3},
	{[]int{-1,-2,-3,0,1}, 2},
	{[]int{-1,2}, 1},
	{[]int{1,2,3,5,-6,4,0,10}, 4},
	{[]int{-10000, -100000}, 2},
	{[]int{5,-20,-20,-39,-5,0,0,0,36,-32,0,-7,-10,-7,21,20,-12,-34,26,2}, 8},
}

func TestGetMaxLen(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if getMaxLen(tcs[i].In) != tcs[i].Expect {
			t.Errorf("get max len test failed on case: %d\n", i)
		}
	}
}
