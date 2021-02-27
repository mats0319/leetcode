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
	{[]int{1,2,3}, [][]int{
		{1,2,3},
		{1,3,2},
		{2,1,3},
		{2,3,1},
		{3,1,2},
		{3,2,1},
	}},
}

func TestPermute(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnTwoDimensionalSlice(permute(tcs[i].In), tcs[i].Expect) {
			t.Errorf("permute test failed on case: %d\n", i)
		}
	}
}
