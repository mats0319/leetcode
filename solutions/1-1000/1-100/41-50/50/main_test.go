package mario

import "testing"

type In struct {
	Float640 float64
	Int0     int
}

var testCase = []struct {
	In     In
	Expect float64
}{
	// test cases here
	{In{2.0, 10}, 1024.0},
	{In{2.1, 3}, 9.261},
	{In{2.0, -2}, 0.25},
}

func TestMyPow(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if myPow(tcs[i].In.Float640, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("my pow test failed on case: %d\n", i)
		}
	}
}
