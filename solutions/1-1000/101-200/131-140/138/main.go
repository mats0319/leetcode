package mario

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	oldList := make([]*Node, 0)
	newList := make([]*Node, 0)

	// make new list with nil random
	res := &Node{}
	{
		p := head
		resP := res
		for p != nil {
			node := &Node{
				Val: p.Val,
			}

			resP.Next = node

			oldList = append(oldList, p)
			newList = append(newList, node)

			p = p.Next
			resP = resP.Next
		}

		res = res.Next
	}

	// maintain random of new list
	p := head
	resP := res
	for p != nil {
		if p.Random != nil {
			resP.Random = newList[getIndex(oldList, p.Random)]
		}

		p = p.Next
		resP = resP.Next
	}

	return res
}

func getIndex(array []*Node, node *Node) int {
	index := -1
	for i := 0; i < len(array); i++ {
		if array[i] == node {
			index = i
			break
		}
	}

	return index
}
