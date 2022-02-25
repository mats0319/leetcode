package mario

type NumArray struct {
	trieSum []int
}

// Constructor nums.length >= 1
func Constructor(nums []int) NumArray {
	trieSum := make([]int, len(nums))
	trieSum[0] = nums[0]
	for i := 1; i < len(trieSum); i++ {
		trieSum[i] = trieSum[i-1] + nums[i]
	}

	return NumArray{
		trieSum: trieSum,
	}
}

func (n *NumArray) Update(index int, val int) {
	diff := 0
	if index == 0 {
		diff = val- n.trieSum[0]
	} else {
		diff = (n.trieSum[index-1] + val) - n.trieSum[index]
	}

    for i := index; i < len(n.trieSum); i++ {
        n.trieSum[i] += diff
    }
}

func (n *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return n.trieSum[right]
    }

    return n.trieSum[right]-n.trieSum[left-1]
}
