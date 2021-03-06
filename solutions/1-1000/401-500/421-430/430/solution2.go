package mario

func flattenTreePreorderTraversal(root *Node) *Node {
	return preorderTraversal(root, &Node{})
}

func preorderTraversal(root *Node, tail *Node) *Node {
	if root == nil {
		return tail
	}

	tail.Next = root
	tail.Next.Prev = tail
	tail.Child = nil
	tail = tail.Next

	tail = preorderTraversal(root.Child, tail)

	tail = preorderTraversal(root.Next, tail)

	return tail
}
