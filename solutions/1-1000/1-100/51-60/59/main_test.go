package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     int
	Expect [][]int
}{
	// test cases here
	{2, [][]int{
		{1,2},
		{4,3},
	}},
}

func TestGenerateMatrix(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnTwoDimensionalSlice(generateMatrix(tcs[i].In), tcs[i].Expect) {
			t.Errorf("generate matrix test failed on case: %d\n", i)
		}
	}
}
