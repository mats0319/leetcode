package mario

import "testing"

var testCase = []struct {
	In     []*ListNode
	Expect *ListNode
}{
	// test cases here
	{[]*ListNode{makeList(7,2,4,3), makeList(5,6,4)}, makeList(7,8,0,7)},
	{[]*ListNode{makeList(5), makeList(5)}, makeList(1,0)},
}

func TestAddTwoNumbers(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnList(addTwoNumbers(tcs[i].In[0], tcs[i].In[1]), tcs[i].Expect) {
			t.Errorf("add two numbers test failed on case: %d\n", i)
		}
	}
}

func makeList(value ...int) *ListNode {
	pre := &ListNode{}
	p := pre
	for i := range value {
		p.Next = &ListNode{Val: value[i]}
		p = p.Next
	}

	return pre.Next
}

func compareOnList(a, b *ListNode) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	for a != nil && b != nil && a.Val == b.Val {
		a = a.Next
		b = b.Next
	}

	return a == nil && b == nil
}
