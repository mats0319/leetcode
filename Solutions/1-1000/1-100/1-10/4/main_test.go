package mario

import (
	"testing"
)

var testCase = []struct {
	In     [][]int
	Expect float64
}{
	{[][]int{{1, 2}, {3, 4}}, 2.5},
	{[][]int{{1, 3}, {2}}, 2.0},
	{[][]int{{1, 4}, {2, 3}}, 2.5},
	{[][]int{{1}, {2, 3}}, 2.0},
}

func TestFindMedianSortedArrays(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findMedianSortedArrays(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("find median sorted arrays test failed on case: %d\n", i)
		}
	}
}
