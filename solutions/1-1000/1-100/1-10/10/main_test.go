package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect bool
}{
	{[]string{"aa", "a"}, false},
	{[]string{"aa", "a*"}, true},
	{[]string{"ab", ".*"}, true},
	{[]string{"aab", "c*a*b"}, true},
	{[]string{"mississippi", "mis*is*p*."}, false},
}

func TestIsMatch(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isMatch(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("regular expression matching test failed on case: %d\n", i)
		}
	}
}

func TestIsMatchDP(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isMatchDP(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("regular expression matching solution: 2 test failed on case: %d\n", i)
		}
	}
}
