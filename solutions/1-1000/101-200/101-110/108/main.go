package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	return toBST(&TreeNode{Val: nums[len(nums)/2]}, nums[:len(nums)/2], nums[len(nums)/2+1:])
}

func toBST(parent *TreeNode, left, right []int) *TreeNode {
	if len(left) > 0 {
		parent.Left = toBST(&TreeNode{Val: left[len(left)/2]}, left[:len(left)/2], left[len(left)/2+1:])
	}

	if len(right) > 0 {
		parent.Right = toBST(&TreeNode{Val: right[len(right)/2]}, right[:len(right)/2], right[len(right)/2+1:])
	}

	return parent
}
