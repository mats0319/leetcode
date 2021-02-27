package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	list := make([]*ListNode, 0)
	var res *ListNode
	for p := head; p.Next != nil; p = p.Next {
		if isContains(list, p) {
			res = p
			break
		}

		list = append(list, p)
	}

	return res
}

func isContains(list []*ListNode, node *ListNode) bool {
	containsFlag := false
	for _, n := range list {
		if n == node {
			containsFlag = true
			break
		}
	}

	return containsFlag
}
