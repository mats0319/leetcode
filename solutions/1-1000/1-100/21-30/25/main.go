package mario

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 {
        return head
    }

    pre := &ListNode{Next: head}
    preNode := pre
    for p := preNode; p != nil; {
        preNodeBack := p.Next

        amount := 0
        for ; amount < k && p != nil; amount++ {
            p = p.Next
        }

        if amount != k || p == nil {
            break
        }

        nextNode := p.Next

        reverseList(preNode, nextNode)

        preNode = preNodeBack
        p = preNode
    }

    return pre.Next
}

func reverseList(preNode, nextNode *ListNode) {
    pre := nextNode
    curr := preNode.Next
    for next := curr.Next; ; {
        curr.Next = pre

        if next == nextNode || next == nil {
            break
        }

        pre = curr
        curr = next
        next = next.Next
    }

    preNode.Next = curr
}
