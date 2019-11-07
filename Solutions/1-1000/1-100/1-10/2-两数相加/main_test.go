package main

import (
	"fmt"
	"testing"
)

var testCase = []struct {
	In     []*ListNode
	Expect *ListNode
}{
	{[]*ListNode{makeList(2, 4, 3), makeList(5, 6, 4)}, makeList(7, 0, 8)},
}

func TestAddTwoNumbers(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnList(addTwoNumbers(tcs[i].In[0], tcs[i].In[1]), tcs[i].Expect) {
			t.Errorf("add two numbers test failed on case: %d", i)
		}
	}
}

func makeList(vals ...int) *ListNode {
	result := new(ListNode)
	p := result
	for i := range vals {
		p.Next = &ListNode{Val: vals[i]}
		p = p.Next
	}

	return result.Next
}

// extra scene :)
func (h *ListNode) printList() {
	p := h
	fmt.Printf("List: ")
	for p.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	fmt.Println(p.Val, " -> nil ")
}

func compareOnList(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}

	return a == nil && b == nil
}
