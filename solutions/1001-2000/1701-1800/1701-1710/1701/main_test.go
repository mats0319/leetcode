package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect float64
}{
	// test cases here
	{[][]int{{1,2}, {2,5}, {4,3}}, 5},
	{[][]int{{5,2}, {5,4}, {10,3}, {20,1}}, 3.25},
}

func TestAverageWaitingTime(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if averageWaitingTime(tcs[i].In) != tcs[i].Expect {
			t.Errorf("average waiting time test failed on case: %d\n", i)
		}
	}
}
