package mario

import (
	"testing"
)

var testCase = []struct {
	In     int
	Expect int
}{
	// test cases here
	{3, 2},
	{6, 9},
}

func TestMinOperations(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minOperations(tcs[i].In) != tcs[i].Expect {
			t.Errorf("min operations test failed on case: %d\n", i)
		}
	}
}
