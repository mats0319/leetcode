package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{2, 3, 2}, 3},
	{[]int{1, 2, 3, 1}, 4},
}

func TestRob(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if rob(tcs[i].In) != tcs[i].Expect {
			t.Errorf("rob test failed on case: %d\n", i)
		}
	}
}
