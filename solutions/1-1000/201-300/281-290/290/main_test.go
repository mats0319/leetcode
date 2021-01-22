package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect bool
}{
	// test cases here
	{[]string{"abba", "dog cat cat dog"}, true},
	{[]string{"abba", "dog cat cat fish"}, false},
	{[]string{"aaaa", "dog cat cat dog"}, false},
	{[]string{"abba", "dog dog dog dog"}, false},
}

func TestWordPattern(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if wordPattern(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("word pattern test failed on case: %d\n", i)
		}
	}
}
