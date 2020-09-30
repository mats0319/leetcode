package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    currVal := head.Val
    p := head.Next
    pre := head
    for p != nil {
        if p.Val == currVal {
            pre.Next = p.Next
        } else {
            currVal = p.Val
            pre = p
        }

        p = p.Next
    }

    return head
}
