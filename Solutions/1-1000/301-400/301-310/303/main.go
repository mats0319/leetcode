package mario

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

type NumArray struct {
	Slice   []int
	Summary []int
}

func Constructor(nums []int) (result NumArray) {
	result.Slice = nums
	result.Summary = make([]int, len(nums)+1)
	for i := 1; i < len(result.Summary); i++ {
		result.Summary[i] = result.Summary[i-1] + result.Slice[i-1]
	}

	return
}

func (n *NumArray) SumRange(i int, j int) int {
	return n.Summary[j+1] - n.Summary[i]
}
