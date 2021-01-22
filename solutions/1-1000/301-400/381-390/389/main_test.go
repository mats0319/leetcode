package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect byte
}{
	// test cases here
	{[]string{"abcd", "abcde"}, 'e'},
	{[]string{"", "y"}, 'y'},
	{[]string{"a", "aa"}, 'a'},
	{[]string{"ae", "aea"}, 'a'},
}

func TestFindTheDifference(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findTheDifference(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("find the difference test failed on case: %d\n", i)
		}
	}
}
