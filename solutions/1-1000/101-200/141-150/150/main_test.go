package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect int
}{
	// test cases here
	{[]string{"2", "1", "+", "3", "*"}, 9},
	{[]string{"4", "13", "5", "/", "+"}, 6},
}

func TestEvalRPN(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if evalRPN(tcs[i].In) != tcs[i].Expect {
			t.Errorf("eval r p n test failed on case: %d\n", i)
		}
	}
}
