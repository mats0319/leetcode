package mario

import "testing"

type In struct {
	IntSlice0 []int
	Int0      int
	Int1      int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In: In{[]int{8, 3}, 6, 5}, Expect: 3},
	{In: In{[]int{10, 9, 6}, 4, 6}, Expect: 7},
	{In: In{[]int{3, 4, 0, 5, 1}, 92, 1}, Expect: -1},
	{In: In{[]int{10, 10,  6, 4, 7}, 8, 3}, Expect: 9},
	{In: In{[]int{0, 4}, 1, 1}, Expect: 2},
}

func TestMinOperationsMaxProfit(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minOperationsMaxProfit(tcs[i].In.IntSlice0, tcs[i].In.Int1, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("min operations max profit test failed on case: %d\n", i)
		}
	}
}
