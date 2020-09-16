package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type ListNode struct {
    Value *TreeNode
    Next  *ListNode
}

// todo: add test file after review tree data-structure
func invertTree(root *TreeNode) *TreeNode {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return root // if root is nil or root is leaf node, return directly
    }

    list := &ListNode{Value: root}
    for ; list != nil; list = list.Next {
        node := list.Value
        node.Left, node.Right = node.Right, node.Left

        p := list
        for p.Next != nil {
            p = p.Next
        }
        if node.Left != nil {
            p.Next = &ListNode{Value: node.Left}
            p = p.Next
        }
        if node.Right != nil {
            p.Next = &ListNode{Value: node.Right}
            p = p.Next
        }
    }

    return root
}
