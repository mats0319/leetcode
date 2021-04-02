package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{10, 3}, 3},
}

func TestDivide(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if divide(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("divide test failed on case: %d\n", i)
		}
	}
}
