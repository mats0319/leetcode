package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }

    res := &ListNode{Next: head}
    pre := res
    p := pre.Next
    for p != nil {
        if p.Val == val {
            pre.Next = p.Next
        } else {
            pre = pre.Next
        }

        p = p.Next
    }

    return res.Next
}
