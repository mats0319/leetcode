package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    m := make(map[*ListNode]bool)

    p := head
    m[p] = true
    hasCycle := false
    for p != nil && !hasCycle {
        p = p.Next

        _, hasCycle = m[p]

        m[p] = true
    }

    return p
}
