package mario

import "math/rand"

type Solution struct {
	data   []int
	length int
}

func Constructor(nums []int) Solution {
	return Solution{
		data:   nums,
		length: len(nums),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.data
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	res := deepCopy(this.data)
	for i := len(res) - 1; i >= 0; i-- {
		index := rand.Intn(i + 1)
		res[index], res[i] = res[i], res[index]
	}

	return res
}

func deepCopy(src []int) []int {
	res := make([]int, len(src))
	for i := 0; i < len(res); i++ {
		res[i] = src[i]
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
