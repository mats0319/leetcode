package mario

import (
	"github.com/mats9693/leetcode/utils/compare"
	"testing"
)

type In struct {
	Is     []int
	Target int
}

var testCase = []struct {
	In     In
	Except [][]int
}{
	{In{[]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0},
		[][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2},
			{-2, 0, 0, 2}, {-1, 0, 0, 1}}},
}

func TestFourSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnErWeiInt(fourSum(tcs[i].In.Is, tcs[i].In.Target), tcs[i].Except) {
			t.Errorf("four sum test failed on case: %d", i)
		}
	}
}
