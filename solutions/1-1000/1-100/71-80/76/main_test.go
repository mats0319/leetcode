package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect string
}{
	// test cases here
	//{[]string{"ADOBECODEBANC", "ABC"}, "BANC"},
	//{[]string{"a", "a"}, "a"},
	//{[]string{"a", "aa"}, ""},
}

func TestMinWindow(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minWindow(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("min window test failed on case: %d\n", i)
		}
	}
}
