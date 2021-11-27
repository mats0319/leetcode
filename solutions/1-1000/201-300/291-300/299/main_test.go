package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect string
}{
	// test cases here
	{[]string{"1122", "1222"}, "3A0B"},
}

func TestGetHint(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if getHint(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("get hint test failed on case: %d\n", i)
		}
	}
}
