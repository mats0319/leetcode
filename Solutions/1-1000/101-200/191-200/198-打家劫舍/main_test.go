package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	{[]int{1, 2, 3, 1}, 4},
	{[]int{2, 7, 9, 3, 1}, 12},
	{nil, 0},
	{[]int{}, 0},
}

func TestRob(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if rob(tcs[i].In) != tcs[i].Expect {
			t.Errorf("rob test failed on case: %d", i)
		}
	}
}
