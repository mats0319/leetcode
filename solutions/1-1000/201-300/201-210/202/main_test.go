package mario

import "testing"

var testCase = []struct {
	In     int
	Expect bool
}{
	// test cases here
	{19, true},
	{2, false},
}

func TestIsHappy(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isHappy(tcs[i].In) != tcs[i].Expect {
			t.Errorf("is happy test failed on case: %d\n", i)
		}
	}
}
