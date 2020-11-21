package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	list := make([]int, 0)
	for p != nil {
	    list = append(list, p.Val)
	    p = p.Next
    }

    QuickSort(list)

	p = head
	for index := 0; p != nil; index++ {
	    p.Val = list[index]
	    p = p.Next
    }

	return head
}

func QuickSort(is []int) {
    if len(is) <= 1 {
        return
    }

    var small int
    {
        var big int
        for i := 1; i < len(is); i++ {
            if is[i] > is[0] {
                big++
            } else {
                small++
                if big != 0 {
                    is[i], is[small] = is[small], is[i]
                }
            }
        }
        if small != 0 {
            is[0], is[small] = is[small], is[0]
        }
    }

    QuickSort(is[:small])
    QuickSort(is[small+1:])

    return
}
