package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// sumOfLeftLeaves root != nil
func sumOfLeftLeaves(root *TreeNode) int {
    return dfs(root, false)
}

func dfs(root *TreeNode, isLeftChild bool) int {
    if root.Left == nil && root.Right == nil {
        value := 0
        if isLeftChild {
            value += root.Val
        }

        return value
    }

    sum := 0

    if root.Left != nil {
        sum += dfs(root.Left, true)
    }

    if root.Right != nil {
        sum += dfs(root.Right, false)
    }

    return sum
}
