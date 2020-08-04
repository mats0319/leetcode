package mario

import "testing"

type In struct {
	Int0      int
	IntSlice0 [][]int
}

var testCase = []struct {
	In     In
	Expect bool
}{
	// test cases here
	{In: In{Int0: 2, IntSlice0: [][]int{{1, 0}}}, Expect: true},
	{In: In{Int0: 2, IntSlice0: [][]int{{1, 0}, {0, 1}}}, Expect: false},
	{In: In{Int0: 2, IntSlice0: [][]int{{0, 1}}}, Expect: true},
}

func TestCanFinish(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if canFinish(tcs[i].In.Int0, tcs[i].In.IntSlice0) != tcs[i].Expect {
			t.Errorf("can finish test failed on case: %d\n", i)
		}
	}
}
