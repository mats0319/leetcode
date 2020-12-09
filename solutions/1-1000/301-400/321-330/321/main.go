package mario

import "net"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
    length1 := len(nums1)
    length2 := len(nums2)

    if k > length1+length2 {
        return nil
    }


}

// make sure k <= slice.length
func maxKDigitWithOrder(slice []int, k int) []int {
    result := make([]int, 0, k)
    maxIndex := 0
    for i := 1; i <= len(slice)-k; i++ {
        if slice[maxIndex] < slice[i] {
            maxIndex = i
        }
    }

    result = append(result, slice[maxIndex])
    result = append(result, maxKDigitWithOrder(slice[maxIndex+1:], k-1)...)

    return result
}

func mergeSlices(a, b []int) []int {
    result := make([]int, 0, len(a) + len(b))
    for i, ia, ib := 0, 0, 0; ia+ib < len(result); {}

}
