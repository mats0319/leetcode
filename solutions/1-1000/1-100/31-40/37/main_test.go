package mario

import "testing"

var testCase = []struct {
	In [][]byte
	Expect [][]byte
}{
	// test cases here
	{},
}

func TestSolveSudoku(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		solveSudoku(tcs[i].In)
		if !compareOnErWeiBytes(tcs[i].In, tcs[i].Expect) {
			t.Errorf("solve sudoku test failed on case: %d\n", i)
		}
	}
}

// make sure a.len = b.len and a[0].len = b[0].len before
func compareOnErWeiBytes(a, b [][]byte) bool {
	isEqual := true
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				isEqual = false
				break
			}
		}
	}

	return isEqual
}
