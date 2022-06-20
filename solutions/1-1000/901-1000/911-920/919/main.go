package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	// deprecated first item, save it layer-order then:
	//   parent index: 'i'
	//   left child index: 2*'i'
	//   right child index: 2*'i'+1
	nodeSlice []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	nodeSlice := make([]*TreeNode, 1) // let length = 1, deprecated first item

	layer := []*TreeNode{root}
	for len(layer) > 0 {
		nextLayer := make([]*TreeNode, 0, len(layer)*2)

		for len(layer) > 0 {
			node := layer[0]
			layer = layer[1:]

			if node == nil { // if 'root' is nil
				break
			}

			nodeSlice = append(nodeSlice, node)

			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		layer = nextLayer
	}

	return CBTInserter{
		nodeSlice: nodeSlice,
	}
}

// Insert return value of parent node
// before insert, tree has at least one node
func (t *CBTInserter) Insert(val int) int {
	newNode := &TreeNode{Val: val}

	// connect parent and new node
	parentIndex := len(t.nodeSlice) / 2

	isLeftChild := len(t.nodeSlice)%2 == 0
	if isLeftChild {
		t.nodeSlice[parentIndex].Left = newNode
	} else {
		t.nodeSlice[parentIndex].Right = newNode
	}

	// append new node at tail of slice
	t.nodeSlice = append(t.nodeSlice, newNode)

	return t.nodeSlice[parentIndex].Val
}

func (t *CBTInserter) Get_root() *TreeNode {
	return t.nodeSlice[1]
}
