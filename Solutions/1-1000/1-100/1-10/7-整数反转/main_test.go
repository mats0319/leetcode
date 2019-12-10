package mario

import "testing"

var testCase = []struct {
	In     int
	Except int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
	{1534236469, 0},
}

func TestReverse(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if reverse(tcs[i].In) != tcs[i].Except {
			t.Errorf("reverse integer test failed on case: %d", i)
		}
	}
}
