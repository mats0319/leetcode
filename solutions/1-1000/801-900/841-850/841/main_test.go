package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect bool
}{
	// test cases here
	{[][]int{{1}, {2}, {3}, {}}, true},
	{[][]int{{1,3},{3,0,1},{2},{0}}, false},
}

func TestCanVisitAllRooms(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if canVisitAllRooms(tcs[i].In) != tcs[i].Expect {
			t.Errorf("can visit all rooms test failed on case: %d\n", i)
		}
	}
}
