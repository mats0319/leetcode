package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{1, 1}, 0},
	{[]int{2, 1}, 0},
	{[]int{2, 2}, 1},
	{[]int{4, 5}, 1},
	// explain:
	// 1: 0
	// 2: 01
	// 3: 0110
	// 4: 01101001
}

func TestKthGrammar(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if kthGrammar(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("kth grammar test failed on case: %d\n", i)
		}
	}
}
