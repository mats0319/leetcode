package mario

import (
	"github.com/mats9693/leetcode/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []int
	Expect [][]int
}{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
	{[]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
		[][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}},
	{[]int{}, [][]int{}},
	{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
}

func TestThreeSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnErWeiInt(threeSum(tcs[i].In), tcs[i].Expect) {
			t.Errorf("three sum test failed on case: %d", i)
		}
	}
}
