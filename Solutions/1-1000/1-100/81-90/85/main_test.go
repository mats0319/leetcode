package mario

import "testing"

var testCase = []struct {
	In     [][]byte
	Expect int
}{
	// test cases here
	{[][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}, 6},
}

func TestMaximalRectangle(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maximalRectangle(tcs[i].In) != tcs[i].Expect {
			t.Errorf("maximal rectangle test failed on case: %d", i)
		}
	}
}
