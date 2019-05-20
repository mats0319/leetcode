package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		head   = &ListNode{Val: 1}
		result = &ListNode{Val: 1}
	)
	makeList(head, 2, 3, 4, 5)
	makeList(result, 2, 3, 5)
	//printList(head)
	//printList(result)
	fmt.Println(compareOnLinkedList(removeNthFromEnd(head, 2), result))
}

func makeList(h *ListNode, vs ...int) {
	p := h
	for i := range vs {
		p.Next = &ListNode{Val: vs[i]}
		p = p.Next
	}
}

func printList(l *ListNode) {
	var t *ListNode
	for t = l; t.Next != nil; t = t.Next {
		fmt.Printf("%d -> ", t.Val)
	}
	fmt.Println(t.Val, ". ")
}

func compareOnList(a, b *ListNode) bool {
	an, bn := a, b
	for an != nil && bn != nil {
		if an.Val != bn.Val {
			return false
		}
		an = an.Next
		bn = bn.Next
	}

	return an == nil && bn == nil
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		step int
		temp *ListNode
	)

	for temp = head; temp != nil; temp = temp.Next {
		step++
	}

	step -= n + 1 // step = step + 1 - n - 2, step is steps need to move.
	if step == -1 {
		return head.Next
	}

	for temp = head; step > 0; step-- {
		temp = temp.Next // temp point to pre of target node finally.
	}
	temp.Next = temp.Next.Next // delete the end is also ok.

	return head
}