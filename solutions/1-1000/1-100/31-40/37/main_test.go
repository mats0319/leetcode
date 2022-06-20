package mario

import "testing"

var testCase = []struct {
	In     [][]byte
	Expect [][]byte
}{
	// test cases here
	{
		In: [][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		Expect: [][]byte{
			{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
			{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
			{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
			{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
			{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
			{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
			{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
			{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
			{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
		},
	},
}

func TestSolveSudoku(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		solveSudoku(tcs[i].In)
		if !compareOnTwoDimensionByteSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("solve sudoku test failed on case: %d\n", i)
		}
	}
}

func compareOnTwoDimensionByteSlice(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := range a {
		if !compareOnByteSlice(a[i], b[i]) {
			isEqual = false
			break
		}
	}

	return isEqual
}

func compareOnByteSlice(a, b []byte) bool {
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
