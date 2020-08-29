package mario

import (
	"testing"
)

var testCase = []struct {
	In     []int
	Expect bool
}{
	// test cases here
	{[]int{2, 6, 4, 1}, false},
	{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true},
}

func TestThreeConsecutiveOdds(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if threeConsecutiveOdds(tcs[i].In) != tcs[i].Expect {
			t.Errorf("three consecutive odds test failed on case: %d\n", i)
		}
	}
}
