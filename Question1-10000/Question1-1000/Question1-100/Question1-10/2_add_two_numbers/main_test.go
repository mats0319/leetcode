package atn

import (
    "fmt"
    "testing"
)

var testCase = [][]*ListNode{
    {makeList(&ListNode{Val: 2}, 4, 3), makeList(&ListNode{Val: 5}, 6, 4)},
}

var result = []*ListNode{
    makeList(&ListNode{Val: 7}, 0, 8),
}

func TestAddTwoNumbers(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if !compareOnList(addTwoNumbers(tcs[i][0], tcs[i][1]), result[i]) {
            t.Errorf("add two numbers test failed on case: %d", i)
        }
    }
}

func makeList(h *ListNode, is ...int) *ListNode {
    p := h
    for i := range is {
        p.Next = &ListNode{Val: is[i]}
        p = p.Next
    }

    return h
}

// extra scene :)
func (h *ListNode) printList() {
    p := h
    for p.Next != nil {
        fmt.Printf("%d -> ", p.Val)
        p = p.Next
    }
    fmt.Println(p.Val, ". ")
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
