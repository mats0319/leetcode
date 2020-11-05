package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

type In struct {
	IntSlice0 [][]int
	IntSlice1 []int
}

var testCase = []struct {
	In     In
	Expect [][]int
}{
	// test cases here
	//{In{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}},
	//	[][]int{{1, 2}, {3, 10}, {12, 16}}},
	//{In{[][]int{}, []int{5, 7}}, [][]int{{5, 7}}},
	//{In{[][]int{{1, 5}}, []int{6, 8}}, [][]int{{1, 5}, {6, 8}}},
	//{In{[][]int{{1, 5}}, []int{0, 0}}, [][]int{{0, 0}, {1, 5}}},
	{In{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, [][]int{{1, 5}, {6, 9}}},
}

func TestInsert(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnTwoDimensionalSlice(insert(tcs[i].In.IntSlice0, tcs[i].In.IntSlice1), tcs[i].Expect) {
			t.Errorf("insert test failed on case: %d\n", i)
		}
	}
}
