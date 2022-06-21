package mario

import "testing"

type In struct {
	Int2Slice0 [][]int
	Bool0      bool
	String0    string
	Int1Slice0 []int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{},
}

func TestTestGenerator(t *testing.T) {
	for i := range testCase {
		if testGenerator(testCase[i].In.String0, testCase[i].In.Int1Slice0, testCase[i].In.Int2Slice0, testCase[i].In.Bool0) != testCase[i].Expect {
			t.Logf("test generator test failed on case: %d\n", i)
			t.Fail()
		}
	}
}
