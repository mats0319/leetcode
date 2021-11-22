package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect string
}{
	// test cases here
	{[]string{"1111", "1111"}, "11110"},
}

func TestAddBinary(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if addBinary(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("add binary test failed on case: %d\n", i)
		}
	}
}
