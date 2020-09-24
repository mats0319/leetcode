package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	inOrder, maxCount := traversalTree(root, make(map[int]int), -1)

	res := make([]int, 0, len(inOrder))
	for k, v := range inOrder {
	    if v == maxCount {
	        res = append(res, k)
        }
    }

	return res
}

func traversalTree(root *TreeNode, inOrder map[int]int, maxCount int) (map[int]int, int) {
	if root == nil {
		return inOrder, maxCount
	}

	inOrder, maxCount = traversalTree(root.Left, inOrder, maxCount)

	inOrder[root.Val]++
	if inOrder[root.Val] > maxCount {
	    maxCount = inOrder[root.Val]
    }

	inOrder, maxCount = traversalTree(root.Right, inOrder, maxCount)

	return inOrder, maxCount
}
