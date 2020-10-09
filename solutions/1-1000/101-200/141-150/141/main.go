package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    m := make(map[*ListNode]bool)

    hasCycleFlag := false
    p := head
    for p != nil && !hasCycleFlag {
        _, hasCycleFlag = m[p]

        m[p] = true
        p = p.Next
    }

    return hasCycleFlag
}
