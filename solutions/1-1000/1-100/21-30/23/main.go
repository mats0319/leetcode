package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}

	var res *ListNode
	for i := 0; i < len(lists); i++ {
		res = mergeList(res, lists[i])
	}

	return res
}

func mergeList(l1, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	res := pre

	for {
		if l1 == nil {
			pre.Next = l2
			break
		}
		if l2 == nil {
			pre.Next = l1
			break
		}

		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}

		pre = pre.Next
	}

	return res.Next
}
