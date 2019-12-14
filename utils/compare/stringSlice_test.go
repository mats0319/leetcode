package cmp

import "testing"

func TestCompareOnStringSlice(t *testing.T) {
	if !CompareOnStringSlice([]string{"ab", "cd", "ef"}, []string{"cd", "ab", "ef"}) {
		t.Errorf("compare on string slice failed.")
	}
}
