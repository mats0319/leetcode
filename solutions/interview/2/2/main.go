package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
    p, q := head, head
    for p != nil && k > 0 {
        p = p.Next
        k--
    }
    for p != nil {
        p = p.Next
        q = q.Next
    }

    return q.Val
}
