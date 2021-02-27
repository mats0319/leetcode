package mario

import "testing"

type In struct {
	IntSlice0 []int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In: In{[]int{1, 2, 3, 4, 5, 6, 1}, 3}, Expect: 12},
	{In: In{[]int{9, 7, 7, 9, 7, 7, 9}, 7}, Expect: 55},
	{In: In{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3}, Expect: 202},
}

func TestMaxScore(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxScore(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("max score test failed on case: %d\n", i)
		}
	}
}
