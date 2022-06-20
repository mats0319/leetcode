package mario

import (
	"testing"
)

var testCase = []struct {
	In     int
	Expect [][]string
}{
	// test cases here
	{4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
}

func TestSolveNQueens(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnTwoDimensionalStringSlice(solveNQueens(tcs[i].In), tcs[i].Expect) {
			t.Errorf("solve n queens test failed on case: %d\n", i)
		}
	}
}

func compareOnTwoDimensionalStringSlice(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := range a {
		if !compareOnStringSlice(a[i], b[i]) {
			isEqual = false
			break
		}
	}

	return isEqual
}

func compareOnStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := range a {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
