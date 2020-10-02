package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    smallPre := &ListNode{}
    s := smallPre
    bigPre := &ListNode{}

    st := smallPre
    bt := bigPre
    for head != nil {
        node := &ListNode{Val: head.Val}
        if node.Val < x {
            st.Next = node
            st = st.Next
            s = s.Next
        } else {
            bt.Next = node
            bt = bt.Next
        }

        head = head.Next
    }

    s.Next = bigPre.Next

    return smallPre.Next
}
