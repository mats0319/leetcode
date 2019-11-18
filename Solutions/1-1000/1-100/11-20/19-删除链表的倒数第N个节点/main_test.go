package main

import (
	"fmt"
	"testing"
)

type In struct {
	Node *ListNode
	N int
}

var testCase = []struct {
	In In
	Except *ListNode
}{
	{In{makeList(1, 2, 3, 4, 5), 2}, makeList(1, 2, 3, 5)},
	{In{makeList(1), 1}, nil},
	{In{makeList(1, 2), 2}, makeList(2)},
}

func TestRemoveNthFromEnd(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnList(removeNthFromEnd(tcs[i].In.Node, tcs[i].In.N), tcs[i].Except) {
			t.Errorf("remove nth from end test failed on case: %d", i)
		}
	}
}

func compareOnList(a, b *ListNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	for a.Next != nil && b.Next != nil {
		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return a.Val == b.Val && a.Next == nil && b.Next == nil
}

func makeList(vals ...int) *ListNode {
	result := &ListNode{}
	p := result

	for i := range vals {
		p.Next = &ListNode{Val: vals[i]}
		p = p.Next
	}

	return result.Next
}

func printList(l *ListNode) {
	p := l

	for l.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	fmt.Printf("%d. ", p.Val)

	return
}
