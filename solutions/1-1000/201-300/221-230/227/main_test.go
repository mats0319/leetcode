package mario

import "testing"

var testCase = []struct {
	In     string
	Expect int
}{
	// test cases here
	{"42", 42},
}

func TestCalculate(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if calculate(tcs[i].In) != tcs[i].Expect {
			t.Errorf("calculate test failed on case: %d\n", i)
		}
	}
}
