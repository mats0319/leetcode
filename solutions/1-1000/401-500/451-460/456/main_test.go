package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect bool
}{
	// test cases here
	{[]int{3, 5, 0, 3, 4}, true},
	{[]int{1, 0, 1, -4, -3}, false},
}

func TestFind132pattern(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if find132pattern(tcs[i].In) != tcs[i].Expect {
			t.Errorf("find132pattern test failed on case: %d\n", i)
		}
	}
}
