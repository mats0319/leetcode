package mario

import "testing"

var testCase = []struct {
	In     *TreeNode
	Expect int
}{
	// test cases here
	{},
}

func TestMinDepth(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minDepth(tcs[i].In) != tcs[i].Expect {
			t.Errorf("min depth test failed on case: %d\n", i)
		}
	}
}
