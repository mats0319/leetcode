package mario

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	layoutTraversal([]*Node{root, nil})

	return root
}

func layoutTraversal(layout []*Node) {
	if len(layout) <= 1 {
		return
	}

	i := 0
	validIndex := len(layout) - 1
	for ; i < validIndex; i++ {
		node := layout[i]
		node.Next = layout[i+1]

		if node.Left != nil {
			layout = append(layout, node.Left)
		}
		if node.Right != nil {
			layout = append(layout, node.Right)
		}
	}

	layout = append(layout, nil)

	layoutTraversal(layout[i+1:])
}
