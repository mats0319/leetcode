package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect bool
}{
	// test cases here
	{[]int{2, -1, 1, 2, 2}, true},
	{[]int{-1, 2}, false},
	{[]int{-2, 1, -1, -2, -2}, false},
	{[]int{-1, -2, -3, -4, -5}, false},
	{[]int{3, 1, 2}, true},
	{[]int{1, 1, 2}, true},
}

func TestCircularArrayLoop(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if circularArrayLoop(tcs[i].In) != tcs[i].Expect {
			t.Errorf("circular array loop test failed on case: %d\n", i)
		}
	}
}
