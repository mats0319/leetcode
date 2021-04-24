package mario

import "testing"

var testCase = []struct {
	In     int
	Expect int
}{
	// test cases here
	{13, 2},
	{4, 1},
}

func TestNumSquares(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if numSquares(tcs[i].In) != tcs[i].Expect {
			t.Errorf("num squares test failed on case: %d\n", i)
		}
	}
}
