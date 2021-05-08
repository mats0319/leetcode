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
	{In{[]int{3,2,2,4,1,4}, 3}, 6},
	{In{[]int{1,2,3,1,1}, 4}, 3},
	{In{[]int{1,2,3,4,5,6,7,8,9,10}, 5}, 15},
}

func TestShipWithinDays(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if shipWithinDays(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("ship within days test failed on case: %d\n", i)
		}
	}
}
