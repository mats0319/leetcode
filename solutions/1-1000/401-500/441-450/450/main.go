package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil && root.Val == key) {
		return nil
	}

	parent := &TreeNode{Left: root}
	parentBackup := parent
	node := root
	isLeftChild := true
	for node != nil && node.Val != key {
		parent = node

		if key < node.Val {
			node = node.Left
			isLeftChild = true
		} else {
			node = node.Right
			isLeftChild = false
		}
	}

	if node == nil { // key is not exist in tree
		return root
	}

	if node.Left != nil && node.Right != nil {
		parent = node
		newNode := node.Left
		isLeftChild = true
		for newNode.Right != nil {
			parent = newNode
			newNode = newNode.Right
			isLeftChild = false
		}

		if isLeftChild {
			parent.Left = newNode.Left
		} else {
			parent.Right = newNode.Left
		}

		node.Val = newNode.Val
	} else {
		if node.Right != nil {
			node = node.Right
		} else { // node.left != nil || (node.left == nil && node.right == nil)
			node = node.Left
		}

		if isLeftChild {
			parent.Left = node
		} else {
			parent.Right = node
		}
	}

	return parentBackup.Left
}
