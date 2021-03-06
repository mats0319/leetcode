package mario

import (
	"testing"
)

type In struct {
	IstNodeSlice0 *ListNode
	Int0          int
}

var testCase = []struct {
	In     In
	Expect *ListNode
}{
	// test cases here
	{In{makeList(1, 2, 3, 4, 5), 2}, makeList(4, 5, 1, 2, 3)},
	{In{makeList(), 2}, makeList()},
	{In{makeList(1, 2), 0}, makeList(1, 2)},
	{In{makeList(1, 2), 1}, makeList(2, 1)},
}

func TestRotateRight(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnList(rotateRight(tcs[i].In.IstNodeSlice0, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("rotate right test failed on case: %d\n", i)
		}
	}
}

func makeList(value ...int) *ListNode {
	list := &ListNode{}
	p := list
	for i := range value {
		p.Next = &ListNode{Val: value[i]}
		p = p.Next
	}

	return list.Next
}

func compareOnList(a, b *ListNode) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	aCopy, bCopy := a, b
	for aCopy != nil && bCopy != nil && a.Val == b.Val {
		aCopy = aCopy.Next
		bCopy = bCopy.Next
	}

	return aCopy == nil && bCopy == nil
}
