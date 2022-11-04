package mario

import "testing"

type In struct {
	FirstVariable  int
	SecondVariable [3]string
	ThirdVariable  [][3]*int
	ForthVariable  bool
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{},
}

func TestTestGenerator1(t *testing.T) {
	for i := range testCase {
		if testGenerator1(testCase[i].In.FirstVariable, testCase[i].In.SecondVariable, testCase[i].In.ThirdVariable, testCase[i].In.ForthVariable) != testCase[i].Expect {
			t.Logf("test generator1 test failed on case: %d\n", i)
			t.Fail()
		}
	}
}
