package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func TestRemoveDuplicates(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if removeDuplicates(tcs[i].In) != tcs[i].Expect {
			t.Errorf("remove duplicates test failed on case: %d\n", i)
		}
	}
}
