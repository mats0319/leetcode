package mario

import "testing"

var testCase = []struct {
	In     string
	Except int
}{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInt(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if romanToInt(tcs[i].In) != tcs[i].Except {
			t.Errorf("roman to int test failed on case: %d", i)
		}
	}
}
