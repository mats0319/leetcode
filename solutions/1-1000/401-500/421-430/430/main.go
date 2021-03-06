package mario

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flattenWithStack(root *Node) *Node {
	if root == nil {
		return root
	}

	stack := make([]*Node, 0)
	stack = append(stack, root)

	listPre := &Node{}
	p := listPre
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for ; ; p.Child = nil {
			p.Next = node
			p.Next.Prev = p
			p = p.Next

			if node.Child == nil {
				if node.Next != nil {
					node = node.Next
					continue
				}

				break
			} else { // if node.Child != nil
				if node.Next != nil {
					stack = append(stack, node.Next)
				}

				node = node.Child
				//continue
			}
		}
	}

	listPre.Next.Prev = nil

	return listPre.Next
}
