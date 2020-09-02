package mario

import "testing"

var testCase = []struct {
	In     string
	Expect bool
}{
	// test cases here
	{"1", true},
}

func TestIsNumber(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isNumber(tcs[i].In) != tcs[i].Expect {
			t.Errorf("is number test failed on case: %d\n", i)
		}
	}
}
