package cmp

import "testing"

func TestCompareOnList(t *testing.T) {
	if !CompareOnList(MakeList(1, 2, 3, 4, 5), MakeList(1, 2, 3, 4, 5)) {
		t.Errorf("compare on list failed.")
	}
}
