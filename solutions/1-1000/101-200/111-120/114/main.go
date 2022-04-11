package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode)  {
    if root == nil {
        return
    }

    nodeList := preOrderTraversal(root)

    p := root
    for i := 1; i < len(nodeList); i++ {
        p.Left = nil
        p.Right = nodeList[i]

        p = p.Right
    }

    p.Left = nil
    p.Right = nil
}

func preOrderTraversal(root *TreeNode) []*TreeNode {
    if root.Left == nil && root.Right == nil {
        return []*TreeNode{root}
    }

    nodeList := []*TreeNode{root}

    if root.Left != nil {
        nodeList = append(nodeList, preOrderTraversal(root.Left)...)
    }

    if root.Right != nil {
        nodeList = append(nodeList, preOrderTraversal(root.Right)...)
    }

    return nodeList
}
