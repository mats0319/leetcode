package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     *cmp.ListNode
	Expect *cmp.ListNode
}{
	// test cases here
	{cmp.MakeList(1, 2, 3, 4, 5), cmp.MakeList(5, 4, 3, 2, 1)},
	{nil, nil},
}

func TestReverseList(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnList(reverseList(tcs[i].In), tcs[i].Expect) {
			t.Errorf("reverse list test failed on case: %d\n", i)
		}
	}
}

//func TestReverseListRecursive(t *testing.T) {
//	tcs := testCase
//	for i := range tcs {
//		if !cmp.CompareOnList(reverseListRecursive(tcs[i].In), tcs[i].Expect) {
//			t.Errorf("reverse list recursive test failed on case: %d\n", i)
//		}
//	}
//}

//func TestReverseListIteration(t *testing.T) {
//	tcs := testCase
//	for i := range tcs {
//		if !cmp.CompareOnList(reverseListIteration(tcs[i].In), tcs[i].Expect) {
//			t.Errorf("reverse list iteration test failed on case: %d\n", i)
//		}
//	}
//}
