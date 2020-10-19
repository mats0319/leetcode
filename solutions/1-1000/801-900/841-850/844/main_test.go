package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect bool
}{
	// test cases here
	{[]string{"ab##", "c#d#"}, true},
	{[]string{"bxj##tw", "bxj###tw"}, false},
}

func TestBackspaceCompare(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if backspaceCompare(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("backspace compare test failed on case: %d\n", i)
		}
	}
}
