package main

import (
	"fmt"
)

// 2584ms, 273MB -- first version
// time least  : 832ms, 310.58% !
// memory least: 161732K
//
// 864ms, 217.3MB -- second version
// 103.85% efficiency, 206.73% advance,
// I will mark the whole process seriously.
func main() {
	fmt.Println(compareErWeiInt(threeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, 0, 1}, {-1, -1, 2}}))
	fmt.Println(compareErWeiInt(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}),
		[][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}))
	fmt.Println(compareErWeiInt(threeSum([]int{}), [][]int{}))
	fmt.Println(compareErWeiInt(threeSum([]int{0, 0, 0}), [][]int{{0, 0, 0}}))
}

func compareErWeiInt(a, b [][]int) bool {
	equal := true
	for {
		if len(a) != len(b) {
			equal = false
			break
		}
		var (
			ma = make(map[int][]int)
			mb = make(map[int][]int)
		)
		for i := 0; i < len(a); i++ {
			ma[i] = a[i]
			mb[i] = b[i]
		}
		var key int
		for i := range ma {
			if key, equal = match(ma[i], mb); !equal {
				break
			}
			delete(ma, i)
			delete(mb, key)
		}

		break
	}

	return equal
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

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort(nums)
	return calculate3sum(nums)
}

func sort(is []int) {
	var maxIndex, minIndex int
	for i := range is {
		if is[i] > is[maxIndex] {
			maxIndex = i
		} else if is[i] < is[minIndex] {
			minIndex = i
		}
	}

	var (
		min  = is[minIndex]
		temp = make([]int, is[maxIndex]-min+1)
	)
	for i := range is {
		temp[is[i]-min]++
	}

	var in int
	for i := range temp {
		for ; temp[i] > 0; temp[i]-- {
			is[in] = i + min
			in++
		}
	}
}

func calculate3sum(is []int) [][]int {
	var (
		left, right, t int
		result         = make([][]int, 0, 100)
	)
	for i := 0; i < len(is)-2; i++ {
		left = i + 1
		right = len(is) - 1
		for left < right {
			t = is[i] + is[left] + is[right]
			if t == 0 {
				result = append(result, []int{is[i], is[left], is[right]})
				for ; left < right && is[left] == is[left+1]; left++ {
				}
				for ; left < right && is[right] == is[right-1]; right-- {
				}
				left++
				right--
			} else if t > 0 {
				right--
			} else {
				left++
			}
		}
		for i < len(is)-2 && is[i] == is[i+1] {
			i++
		}
	}
	return result
}

//func threeSum(nums []int) [][]int {
//	var (
//		zeroNum, ip, in int
//		positive        = make([]int, 0, len(nums))
//		negative        = make([]int, 0, len(nums))
//		result          = make([][]int, 0)
//	)
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == 0 {
//			zeroNum++
//		} else if nums[i] > 0 {
//			positive = append(positive, nums[i])
//			ip++
//		} else {
//			negative = append(negative, nums[i])
//			in++
//		}
//	}
//	positive = sortAndThrowMoreThanTwoRepeat(positive[:ip])
//	negative = sortAndThrowMoreThanTwoRepeat(negative[:in])
//
//	if zeroNum >= 3 {
//		result = append(result, []int{0, 0, 0})
//	}
//	if zeroNum > 0 {
//		result = containOneZero(positive, negative, result)
//	}
//
//	result = noZero(positive, negative, result, true)
//	result = noZero(negative, positive, result, false)
//
//	fmt.Println(positive, negative, result)
//
//	return result
//}
//
//func sortAndThrowMoreThanTwoRepeat(is []int) []int {
//	var minIndex int
//	for i := 0; i < len(is); i++ {
//		minIndex = i
//		for j := i + 1; j < len(is); j++ {
//			if is[j] < is[minIndex] {
//				minIndex = j
//			}
//		}
//		if minIndex != i {
//			is[i], is[minIndex] = is[minIndex], is[i]
//		}
//	}
//
//	result := make([]int, 0, len(is))
//	var mark int
//	for i := 0; i < len(is)-2; i++ {
//		if is[i] == is[i+1] && is[i+1] == is[i+2] {
//			result = append(result, is[mark:i]...)
//			mark = i + 1
//		}
//	}
//	result = append(result, is[mark:]...)
//	return result
//}
//
//func containOneZero(src, dst []int, result [][]int) [][]int {
//	for is := range src {
//		if is < len(src)-1 && src[is] == src[is+1] {
//			continue
//		}
//		for id := range dst {
//			if dst[id] == -src[is] {
//				result = append(result, []int{dst[id], 0, src[is]})
//				break
//			}
//		}
//	}
//	return result
//}
//
//func noZero(src, dst []int, result [][]int, srcPos bool) [][]int {
//	var t int
//	for is := range src {
//		if is < len(src)-1 && src[is] == src[is+1] {
//			continue
//		}
//		for i1 := 0; i1 < len(dst)-1; i1++ {
//			t = src[is] + dst[i1]
//			for i2 := i1 + 1; i2 < len(dst); i2++ {
//				if dst[i2] == -t {
//					if srcPos {
//						result = append(result, []int{dst[i1], dst[i2], src[is]})
//						for i1 < len(dst)-1 && dst[i1] == dst[i1+1] {
//							i1++
//						}
//					} else {
//						result = append(result, []int{src[is], dst[i1], dst[i2]})
//						for i1 < len(dst)-1 && dst[i1] == dst[i1+1] {
//							i1++
//						}
//					}
//					break
//				}
//			}
//		}
//	}
//	return result
//}
