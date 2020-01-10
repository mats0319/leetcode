package mario

import (
	"github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []int
	Expect [][]int
}{
	{[]int{0, -3, -2, -1, 0, 0, 1, 2, 3},
		[][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2},
			{-2, 0, 0, 2}, {-1, 0, 0, 1}}},
}

func TestFourSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnErWeiInt(fourSum(tcs[i].In[1:], tcs[i].In[0]), tcs[i].Expect) {
			t.Errorf("four sum test failed on case: %d\n", i)
		}
	}
}
