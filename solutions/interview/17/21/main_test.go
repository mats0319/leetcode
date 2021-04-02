package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
}

func TestTrap(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if trap(tcs[i].In) != tcs[i].Expect {
			t.Errorf("trap test failed on case: %d\n", i)
		}
	}
}
