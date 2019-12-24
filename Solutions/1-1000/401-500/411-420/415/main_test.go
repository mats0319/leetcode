package mario

import (
	"testing"
)

var testCase = []struct {
	In     []string
	Expect string
}{
	{[]string{"123", "4567"}, "4690"},
	{[]string{"0", "0"}, "0"},
	{[]string{"9", "99"}, "108"},
	{[]string{"584", "18"}, "602"},
}

func TestAddStrings(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if addStrings(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("add strings test failed on case: %d", i)
		}
	}
}
