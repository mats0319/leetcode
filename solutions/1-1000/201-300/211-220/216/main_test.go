package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []int
	Expect [][]int
}{
	// test cases here
	{[]int{3, 7}, [][]int{{1, 2, 4}}},
	{[]int{3, 9}, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	{[]int{3, 2}, [][]int{}},
	{[]int{4, 24}, [][]int{{1, 6, 8, 9}, {2, 5, 8, 9}, {2, 6, 7, 9}, {3, 4, 8, 9}, {3, 5, 7, 9}, {3, 6, 7, 8}, {4, 5, 6, 9}, {4, 5, 7, 8}}},
}

func TestCombinationSum3(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		res = make([][]int, 0)
		if !cmp.CompareOnTwoDimensionalSlice(combinationSum3(tcs[i].In[0], tcs[i].In[1]), tcs[i].Expect) {
			t.Errorf("combination sum3 test failed on case: %d\n", i)
		}
	}
}
