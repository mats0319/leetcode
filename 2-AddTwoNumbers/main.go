package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		l1 *ListNode = new(ListNode)
		l2 *ListNode = new(ListNode)
		t  *ListNode
	)
	l1.Val = 2
	t = new(ListNode)
	t.Val = 4
	l1.Next = t
	t = new(ListNode)
	t.Val = 3
	l1.Next.Next = t
	l2.Val = 5
	t = new(ListNode)
	t.Val = 6
	l2.Next = t
	t = new(ListNode)
	t.Val = 4
	l2.Next.Next = t
	lr := addTwoNumbers(l1, l2)
	for {
		fmt.Printf("%+v", l1.Val)
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
		fmt.Printf(" -> ")
	}
	fmt.Println(".")
	for {
		fmt.Printf("%+v", l2.Val)
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
		fmt.Printf(" -> ")
	}
	fmt.Println(".")
	for {
		fmt.Printf("%+v", lr.Val)
		if lr.Next == nil {
			break
		}
		lr = lr.Next
		fmt.Printf(" -> ")
	}
	fmt.Println(".")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lr := new(ListNode)
	p := lr

	var flag, sum int

	for {
		sum = flag
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Val = sum % 10
		flag = sum / 10
		
		if l1 == nil && l2 == nil && flag == 0 {
			break
		}
		p.Next = new(ListNode)
		p = p.Next
	}

	return lr
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var lr *ListNode = new(ListNode)
	p := lr

	var sum, flag int
	for {
		sum = l1.Val + l2.Val + flag
		p.Val = sum % 10
		flag = sum / 10
		p.Next = new(ListNode)
		if l1.Next == nil || l2.Next == nil {
			if flag == 1 {
				p.Next.Val = 1
			} else {
				p.Next = nil
			}
			break
		}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1.Next == nil && l2.Next == nil {
		return lr
	}

	sum = 0
	p.Next = new(ListNode)
	p = p.Next
	if l1.Next != nil {
		l1 = l1.Next
		for {
			sum = l1.Val + flag
			p.Val = sum % 10
			flag = sum / 10
			if l1.Next == nil {
				if flag == 1 {
					p.Next = new(ListNode)
					p.Next.Val = 1
				}
				break
			}
			p.Next = new(ListNode)
			p = p.Next
			l1 = l1.Next
		}
	}
	if l2.Next != nil {
		l2 = l2.Next
		for {
			sum = l2.Val + flag
			p.Val = sum % 10
			flag = sum / 10
			if l2.Next == nil {
				if flag == 1 {
					p.Next = new(ListNode)
					p.Next.Val = 1
				}
				break
			}
			p.Next = new(ListNode)
			p = p.Next
			l2 = l2.Next
		}
	}

	return lr
}
*/