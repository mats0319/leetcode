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
	{
		[]int{4, 4, 4, 1, 4}, [][]int{
			{},
			{4},
			{4, 4},
			{4, 4, 4},
			{1},
			{4, 1},
			{4, 4, 1},
			{4, 4, 4, 1},
			{4, 4, 4, 4},
			{4, 4, 4, 1, 4},
		},
	},
	{
		[]int{1, 2, 2}, [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{2, 2},
			{1, 2, 2},
		},
	},
	{
		[]int{0}, [][]int{
			{},
			{0},
		},
	},
}

func TestSubsetsWithDup(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnTwoDimensionalSlice(subsetsWithDup(tcs[i].In), tcs[i].Expect) {
			t.Errorf("subsets with dup test failed on case: %d\n", i)
		}
	}
}
