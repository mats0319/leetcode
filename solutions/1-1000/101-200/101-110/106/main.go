package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) < 1 {
        return nil
    }

    index := findIndex(inorder, postorder[len(postorder)-1])

    node := &TreeNode{Val: postorder[len(postorder)-1]}
    node.Left = buildTree(inorder[:index], postorder[:index])
    node.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

    return node
}

func findIndex(array []int, value int) int {
    index := 0
    for index < len(array) && array[index] != value {
        index++
    }

    return index
}
