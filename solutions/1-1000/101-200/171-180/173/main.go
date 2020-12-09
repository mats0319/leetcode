package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	increase  []int
	nextIndex int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		increase:  inorderTraversal(root, make([]int, 0)),
		nextIndex: 0,
	}
}

func inorderTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = inorderTraversal(root.Left, res)

	res = append(res, root.Val)

	res = inorderTraversal(root.Right, res)

	return res
}

func (this *BSTIterator) Next() int {
	res := this.increase[this.nextIndex]
	this.nextIndex++

	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.nextIndex < len(this.increase)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
