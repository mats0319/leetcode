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

	traversalSequence([]*Node{root, nil})

	return root
}

func traversalSequence(currLayerNodes []*Node) {
	if len(currLayerNodes) < 2 {
		return
	}

	nextLayerNodes := make([]*Node, 0)
	for i := 0; i < len(currLayerNodes)-1; i++ {
		node := currLayerNodes[i]
		node.Next = currLayerNodes[i+1]

		if node.Left != nil {
			nextLayerNodes = append(nextLayerNodes, node.Left)
		}
		if node.Right != nil {
			nextLayerNodes = append(nextLayerNodes, node.Right)
		}
	}

	traversalSequence(append(nextLayerNodes, nil))
}
