package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode)  {
    stack := make([]*ListNode, 0)
    for p := head; p != nil; p = p.Next {
        stack = append(stack, p)
    }

    l, r := 0, len(stack)-1
    pre := &ListNode{}
    p := pre
    for l < r {
        p.Next = stack[l]
        p = p.Next
        p.Next = stack[r]
        p = p.Next

        l++
        r--
    }

    if l == r {
        p.Next = stack[l]
        p = p.Next
    }

    p.Next = nil

    head = pre.Next

    return
}
