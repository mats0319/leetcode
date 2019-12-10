package cmp

import "fmt"

// ListNode define node structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// CompareOnList compare if two list is equal
func CompareOnList(a, b *ListNode) bool {
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

// MakeList make a list
func MakeList(vals ...int) *ListNode {
	result := &ListNode{}
	p := result

	for i := range vals {
		p.Next = &ListNode{Val: vals[i]}
		p = p.Next
	}

	return result.Next
}

// PrintList print a list
func PrintList(l *ListNode) {
	p := l

	for l.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	fmt.Printf("%d. ", p.Val)

	return
}
