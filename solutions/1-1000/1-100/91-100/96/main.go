package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numTrees(n int) int {
	data := make([]int, n)
	for i := range data {
		data[i] = i + 1
	}

	return buildAllBST(data)
}

func buildAllBST(v []int) int {
	if len(v) < 3 {
		return len(v)
	}

	res := 0
	for i := range v {
		left := buildAllBST(v[:i])
		if left == 0 {
			left = 1
		}

		right := buildAllBST(v[i+1:])
		if right == 0 {
			right = 1
		}

		res += left * right
	}

	return res
}
