package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect int
}{
	{[]string{"horse", "ros"}, 3},
	{[]string{"intention", "execution"}, 5},
}

func TestMinDistance(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minDistance(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("min distance test failed on case: %d", i)
		}
	}
}
