package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	{[]int{1, -1, 2, 1, -4}, 2},
	{[]int{3, 0, 1, 2}, 3},
	{[]int{-100, 1, 1, 1, 0}, 2},
	{[]int{-1, 1, 1, -1, -1, 3}, -1},
	{[]int{-1, -3, -2, -5, 3, -4}, -2},
	{[]int{1, 0, 2, 1, -3}, 0},
}

func TestThreeSumClosest(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if threeSumClosest(tcs[i].In[1:], tcs[i].In[0]) != tcs[i].Expect {
			t.Errorf("three sum closest test failed on case: %d", i)
		}
	}
}
