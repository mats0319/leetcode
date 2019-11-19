package main

import (
	"github.com/mats9693/leetcode/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []*compare.ListNode
	Expect *compare.ListNode
}{
	{[]*compare.ListNode{compare.MakeList(2, 4, 3), compare.MakeList(5, 6, 4)}, compare.MakeList(7, 0, 8)},
}

func TestAddTwoNumbers(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compare.CompareOnList(addTwoNumbers(tcs[i].In[0], tcs[i].In[1]), tcs[i].Expect) {
			t.Errorf("add two numbers test failed on case: %d", i)
		}
	}
}
