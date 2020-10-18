package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    odd, even := head, head.Next
    evenBackup := even
    for even != nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }

    odd.Next = evenBackup

    return head
}
