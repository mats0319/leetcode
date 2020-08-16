package mario

import (
	"testing"
)

var testCase = []struct {
	In     int
	Expect int
}{
	// test cases here
	{10, 4},
	{6, 3},
	{1, 1},
	{56, 6},
	{182, 8},
	//{84806671, 0}, // panic?
}

func TestMinDays(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minDays(tcs[i].In) != tcs[i].Expect {
			t.Errorf("min days test failed on case: %d\n", i)
		}
	}
}
