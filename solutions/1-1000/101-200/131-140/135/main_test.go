package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{1, 0, 2}, 5},
	{[]int{1, 2, 2}, 4},
	{[]int{1, 3, 2, 2, 1}, 7},
}

func TestCandy(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if candy(tcs[i].In) != tcs[i].Expect {
			t.Errorf("candy test failed on case: %d\n", i)
		}
	}
}
