package mario

import "testing"

type In struct {
	IntSlice0 [][]int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{[][]int{{0,2}, {4,6}, {8, 10}, {1, 9}, {1, 5}, {5,9}}, 10}, 3},
	{In{[][]int{{0,1}, {1,2}}, 5}, -1},
	{In{[][]int{{0,1}, {6,8}, {0, 2}, {5, 6}, {0, 4}, {0,3}, {6,7}, {1,3}, {4,7},
		{1,4}, {2,5}, {2,6}, {3,4}, {4,5}, {5,7}, {6,9}}, 9}, 3},
	{In{[][]int{{0,2}, {4,8}}, 5}, -1},
}

func TestVideoStitching(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if videoStitching(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("video stitching test failed on case: %d\n", i)
		}
	}
}
