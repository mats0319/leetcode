package mario

import (
	"fmt"
	"testing"
)

var testCase = []struct {
	In     string
	Expect int
}{
	// test cases here
	{"(1+(4+5+2)-3)+(6+8)", 23},
	{"2147483647", 2147483647},
}

func TestCalculate(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if calculate(tcs[i].In) != tcs[i].Expect {
			t.Errorf("calculate test failed on case: %d\n", i)
			fmt.Println(calculate(tcs[i].In), tcs[i].Expect)
		}
	}
}
