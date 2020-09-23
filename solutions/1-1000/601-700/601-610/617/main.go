package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    return NewTree(t1, t2, &TreeNode{})
}

func NewTree(t1, t2, root *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    }
    if t1 == nil {
        t1 = &TreeNode{}
    }
    if t2 == nil {
        t2 = &TreeNode{}
    }

    root = &TreeNode{Val: t1.Val + t2.Val}

    root.Left = NewTree(t1.Left, t2.Left, root.Left)

    root.Right = NewTree(t1.Right, t2.Right, root.Right)

    return root
}
