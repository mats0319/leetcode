package cmp

import "testing"

func TestCompareOnErWeiInt(t *testing.T) {
	if !CompareOnErWeiInt([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 8, 9}, {1, 2, 3}, {6, 5, 4}}) {
		t.Errorf("compare on er-wei int failed.")
	}
}
