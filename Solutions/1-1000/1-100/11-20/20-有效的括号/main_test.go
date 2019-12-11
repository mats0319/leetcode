package mario

import "testing"

var testCase = []struct {
	In     string
	Except bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
	{")(", false},
}

func TestIsValid(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isValid(tcs[i].In) != tcs[i].Except {
			t.Errorf("is valid test failed on case: %d", i)
		}
	}
}
