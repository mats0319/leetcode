package mario

// maxChunksToSorted
func maxChunksToSorted(arr []int) int {
    // max chunks means as small as possible for each chunk,
    // a method is loop: from 'index' to its left is a chunk,
    // like '[1,3,0,2,4]', first chunk must include '0',
    // and base on it, '0' value is on '2' index now, we need to find if values on index '0~2' are just '0~2',
    // if there has an exceed value, we need to extend current chunk
    res := 0
    for i := 0; i < len(arr); i++ {
        maxIndex := i
        for maxIndex < len(arr) && arr[maxIndex] != i {
            maxIndex++
        }

        for j := i; j <= maxIndex; j++ {
            if arr[j] > maxIndex {
                maxIndex = arr[j]
            }
        }

        i = maxIndex
        res++
    }

    return res
}
