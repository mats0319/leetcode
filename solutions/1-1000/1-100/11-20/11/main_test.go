package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
}

func TestMaxArea(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxArea(tcs[i].In) != tcs[i].Expect {
			t.Errorf("container with most water test failed on case: %d\n", i)
		}
	}
}
