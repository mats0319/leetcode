package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1, l2 = convertList(l1), convertList(l2)

    pre := &ListNode{}
    lr := pre
    addFlag := 0
    for l1 != nil || l2 != nil || addFlag == 1 {
        v := getValue(l1)+getValue(l2)+addFlag
        addFlag = v / 10
        v %= 10

        lr.Next = &ListNode{Val: v}
        lr = lr.Next

        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }

    return convertList(pre.Next)
}

func convertList(l *ListNode) *ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    pre := (*ListNode)(nil)
    curr := l
    for curr != nil {
        next := curr.Next

        curr.Next = pre
        pre = curr
        curr = next
    }

    return pre
}

func getValue(node *ListNode) int {
    res := 0
    if node != nil {
        res = node.Val
    }

    return res
}
