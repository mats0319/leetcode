package mario

import "testing"

var testCase = []struct {
	In     [][]byte
	Expect int
}{
	// test cases here
	{[][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}, 6},
	{[][]byte{{'1'}}, 1},
	{[][]byte{[]byte("10")}, 1},
	{[][]byte{[]byte("01")}, 1},
	{[][]byte{
		[]byte("01"),
		[]byte("01"),
	}, 2},
}

func TestMaximalRectangle(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maximalRectangle(tcs[i].In) != tcs[i].Expect {
			t.Errorf("maximal rectangle test failed on case: %d\n", i)
		}
	}
}
