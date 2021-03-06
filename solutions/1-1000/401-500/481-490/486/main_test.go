package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect bool
}{
	// test cases here
	{[]int{1, 5, 2}, false},
	{[]int{1, 5, 233, 7}, true},
	{[]int{0}, true},
	{[]int{1, 5, 2, 4, 6}, true},
}

func TestPredictTheWinner(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if PredictTheWinner(tcs[i].In) != tcs[i].Expect {
			t.Errorf("predict the winner test failed on case: %d\n", i)
		}
	}
}
