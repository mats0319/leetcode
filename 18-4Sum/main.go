package main

import (
	"fmt"
)

func main() {
	//fmt.Println(compareErWeiInt(fourSum([]int{1, 0, -1, 0, -2, 2}, 0),
	//	[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}))
	//fmt.Println(compareErWeiInt(fourSum([]int{2, 1, 0, -1}, 2), [][]int{{2, 1, 0, -1}}))
	fmt.Println(compareErWeiInt(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0),
		[][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}))
}

func compareErWeiInt(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	var (
		ma = make(map[int][]int)
		mb = make(map[int][]int)
	)
	for i := 0; i < len(a); i++ {
		ma[i] = a[i]
		mb[i] = b[i]
	}

	var (
		key   int
		equal bool
	)
	for i := range ma {
		if key, equal = match(ma[i], mb); !equal {
			return false
		}
		delete(ma, i)
		delete(mb, key)
	}

	return true
}
func match(a []int, b map[int][]int) (int, bool) {
	var (
		key     int
		isMatch bool
	)
	for key = range b {
		isMatch = true
		for i := 0; i < 3; i++ {
			if b[key][i] != a[i] {
				isMatch = false
				break
			}
		}
		if isMatch {
			break
		}
	}
	return key, isMatch
}

// dynamic programming algorithm attempt!
// sort, we have a condition: a <= b <= c <= d to implement a + b + c + d = t
// a = 0       , a < length - n + 1   (n = 4)
// b = a + 1   , b < length - n + 2
// |- c = b + 1, c < length - n + 3 && c < d
// |- d = length - 1
// c++, d-- and make table dp[c][d] = arr[c] + arr[d]
// only iterate array twice after make table, but make table may cost a lot of time?
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	if len(nums) == 4 {
		if sum(nums) == target {
			return [][]int{nums}
		} else {
			return nil
		}
	}

	sort(nums)

	var (
		a, b   int
		length = len(nums)
		c, d   = 2, length - 1
		dp     = make([][]int, length)
	)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for ; c < d; c++ {
		for ; d > c; d-- {
			dp[c][d] = nums[c] + nums[d]
		}
		d = length - 1
	}
	fmt.Println("dp: ", dp)

	var (
		temp1, temp2 int
		result       = make([][]int, 0, 20)
	)
	for ; a < length-3; a++ {
		temp1 = target - nums[a]
		for b = a + 1; b < length-2; b++ {
			temp2 = temp1 - nums[b]
			for c = b + 1; c < d; c++ {
				for ; d > c; d-- {
					if dp[c][d] == temp2 {
						result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})
						break
					}
				}
				d = length - 1
				for c < d && nums[c] == nums[c+1] {
					c++
				}
			}
			for b < length-2 && nums[b] == nums[b+1] {
				b++
			}
		}
		for a < length-3 && nums[a] == nums[a+1] {
			a++
		}
	}

	return result
}

func sum(is []int) (sum int) {
	for i := range is {
		sum += is[i]
	}

	return
}

func sort(is []int) {
	for i := len(is) - 1; i >= 0; i-- {
		var maxIndex int
		for j := 0; j < i+1; j++ {
			if is[j] > is[maxIndex] {
				maxIndex = j
			}
		}
		is[i], is[maxIndex] = is[maxIndex], is[i]
	}
}
