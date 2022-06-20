package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	// test cases here
	{},
}

func TestTestGenerator(t *testing.T) {
	for i := range testCase {
		if testGenerator(testCase[i].In[0], testCase[i].In[1]) != testCase[i].Expect {
			t.Logf("test generator test failed on case: %d\n", i)
			t.Fail()
		}
	}
}
