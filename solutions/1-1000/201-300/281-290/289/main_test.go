package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     [][]int
	Expect [][]int
}{
	// test cases here
	{In: [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}, Expect: [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}},
}

func TestGameOfLife(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		gameOfLife(tcs[i].In)
		if !cmp.CompareOnTwoDimensionalSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("game of life test failed on case: %d\n", i)
		}
	}
}
