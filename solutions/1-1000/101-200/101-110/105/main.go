package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) < 1 {
        return nil
    }

    index := findIndex(inorder, preorder[0])

    node := &TreeNode{Val: preorder[0]}
    node.Left = buildTree(preorder[1:index+1], inorder[:index])
    node.Right = buildTree(preorder[index+1:], inorder[index+1:])

    return node
}

func findIndex(array []int, value int) int {
    index := 0
    for index < len(array) && array[index] != value {
        index++
    }

    return index
}
