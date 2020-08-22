package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect bool
}{
	// test cases here
	{},
}

func TestJudgePoint24(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if judgePoint24(tcs[i].In) != tcs[i].Expect {
			t.Errorf("judge point24 test failed on case: %d\n", i)
		}
	}
}
