package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect bool
}{
	{[]string{"aa", "a"}, false},
	{[]string{"aa", "*"}, true},
	{[]string{"cb", "?a"}, false},
	{[]string{"adceb", "*a*b"}, true},
	{[]string{"acdcb", "a*c?b"}, false},
}

func TestIsMatch(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isMatch(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("is match test failed on case: %d\n", i)
		}
	}
}
