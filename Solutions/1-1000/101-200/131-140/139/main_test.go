package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect bool
}{
	{[]string{"leetcode", "leet", "code"}, true},
	{[]string{"applepenapple", "apple", "pen"}, true},
	{[]string{"catsandog", "cats", "dog", "sand", "and", "cat"}, false},
}

func TestWordBreak(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if wordBreak(tcs[i].In[0], tcs[i].In[1:]) != tcs[i].Expect {
			t.Errorf("word break test failed on case: %d\n", i)
		}
	}
}
