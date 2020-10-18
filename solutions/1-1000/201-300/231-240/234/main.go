package mario

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }

    payload := make([]int, 0)
    for p := head; p != nil; p = p.Next {
        payload = append(payload, p.Val)
    }

    start, end := 0, len(payload)-1
    for start < end {
        if payload[start] != payload[end] {
            break
        }

        start++
        end--
    }

    return start >= end
}
