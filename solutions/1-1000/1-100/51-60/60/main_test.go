package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect string
}{
	// test cases here
	{In: []int{3, 3}, Expect: "213"},
}

func TestGetPermutation(t *testing.T) {
	for i := range testCase {
		if getPermutation(testCase[i].In[0], testCase[i].In[1]) != testCase[i].Expect {
			t.Logf("get permutation test failed on case: %d\n", i)
			t.Fail()
		}
	}
}
