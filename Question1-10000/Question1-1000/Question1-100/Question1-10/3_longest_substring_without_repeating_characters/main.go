package lswrc

func lengthOfLongestSubstring(s string) int {
    if len(s) <= 1 {
        return len(s)
    }

    var indexArray [256]int
    for i := 0; i < 256; i++ {
        indexArray[i] = -1
    }

    var result, length, start int
    for i := range s {
        if indexArray[s[i]] != -1 {
            if indexArray[s[i]] >= start && length > result {
                result = length
            }
            if indexArray[s[i]] + 1 > start {
                start = indexArray[s[i]] + 1
            }
            length = i - start
        }
        indexArray[s[i]] = i
        length++
    }

    if length > result {
        result = length
    }

    return result
}
