package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect bool
}{
	// test cases here
	{[]string{"vtkgn", "vttkgnn"}, true},
}

func TestIsLongPressedName(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isLongPressedName(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("is long pressed name test failed on case: %d\n", i)
		}
	}
}
