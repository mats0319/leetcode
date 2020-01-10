package mario

import "testing"

var testCase = []struct {
	In     [][]byte
	Expect int
}{
	{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 4},
	{[][]byte{}, 0},
	{[][]byte{{'0'}}, 0},
	{[][]byte{{'0', '1'}}, 1},
}

func TestMaximalSquare(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maximalSquare(tcs[i].In) != tcs[i].Expect {
			t.Errorf("maximal square test failed on case: %d\n", i)
		}
	}
}
