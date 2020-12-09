package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func countNodes(root *TreeNode) int {
    return preorderTraversal(root, 0)
}

func preorderTraversal(root *TreeNode, count int) int {
    if root == nil {
        return count
    }

    count++

    count = preorderTraversal(root.Left, count)

    count = preorderTraversal(root.Right, count)

    return count
}
