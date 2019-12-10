package mario

import "testing"

type In struct {
	Is     []int
	Target int
}

var testCase = []struct {
	In     In
	Except int
}{
	{In{[]int{-1, 2, 1, -4}, 1}, 2},
	{In{[]int{0, 1, 2}, 3}, 3},
	{In{[]int{1, 1, 1, 0}, -100}, 2},
	{In{[]int{1, 1, -1, -1, 3}, -1}, -1},
	{In{[]int{-3, -2, -5, 3, -4}, -1}, -2},
	{In{[]int{0, 2, 1, -3}, 1}, 0},
}

func TestThreeSumClosest(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if threeSumClosest(tcs[i].In.Is, tcs[i].In.Target) != tcs[i].Except {
			t.Errorf("three sum closest test failed on case: %d", i)
		}
	}
}
