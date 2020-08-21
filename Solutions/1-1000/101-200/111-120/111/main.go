package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) (depth int) {
    nodeSlice := make([]*TreeNode, 0)
    depthSlice := make([]int, 0)

    nodeSlice = append(nodeSlice, root)
    depthSlice = append(depthSlice, 1)
    for i := 0; i < len(nodeSlice); i++ {
        n := nodeSlice[i]

        if n == nil {
            continue
        }

        depth = depthSlice[i]

        if n.Left == nil && n.Right == nil {
            break
        }

        if n.Left != nil {
            nodeSlice = append(nodeSlice, n.Left)
            depthSlice = append(depthSlice, depth+1)
        } else if n.Right != nil {
            nodeSlice = append(nodeSlice, n.Right)
            depthSlice = append(depthSlice, depth+1)
        }
    }

    return
}
