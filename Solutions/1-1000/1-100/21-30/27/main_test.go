package mario

import "testing"

type In struct {
	Int0      int
	IntSlice0 []int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{3, []int{3,2,2,3}}, 2},
	{In{2, []int{0,1,2,2,3,0,4,2}}, 5},
}

func TestRemoveElement(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if removeElement(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("remove element test failed on case: %d", i)
		}
	}
}
