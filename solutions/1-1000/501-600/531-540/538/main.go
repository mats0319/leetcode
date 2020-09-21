package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    postOrderTraversal(root, 0)

    return root
}

func postOrderTraversal(root *TreeNode, sum int) int {
    if root == nil {
        return sum
    }

    sum = postOrderTraversal(root.Right, sum)

    sum += root.Val
    root.Val = sum

    sum = postOrderTraversal(root.Left, sum)

    return sum
}
