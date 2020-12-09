package mario

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }

    values := make([]int, 0)
    for head != nil {
        values = append(values, head.Val)
        head = head.Next
    }

    return makeTree(values)
}

func makeTree(values []int) *TreeNode {
    split := 0
    {
        length := len(values)

        if length < 1 {
            return nil
        } else if length < 2 {
            return &TreeNode{Val: values[0]}
        }

        split = length / 2
    }

    root := &TreeNode{Val: values[split]}

    root.Left = makeTree(values[:split])
    root.Right = makeTree(values[split+1:])

    return root
}
