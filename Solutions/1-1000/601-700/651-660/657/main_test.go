package mario

import "testing"

var testCase = []struct {
	In     string
	Expect bool
}{
	// test cases here
	{"UD", true},
	{"LL", false},
}

func TestJudgeCircle(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if judgeCircle(tcs[i].In) != tcs[i].Expect {
			t.Errorf("judge circle test failed on case: %d\n", i)
		}
	}
}
