package mario

import "sort"

// s, t not null
// todo: debug
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// format t into [52]int
	// traversal s, record each target char's index (target char: char in t), get first start-end pointer
	// maintain start-end pointer, until one char empty
	formatT := make([]int, 52)
	for i := range t {
		formatT[byteToIndex(t[i])]++
	}

	validSIndexForChars := make([][]int, 52) // char index - indexes in s
	validSIndexes := make([]int, 0)
	for i := range s {
		if formatT[byteToIndex(s[i])] > 0 {
			indexSlice := validSIndexForChars[byteToIndex(s[i])]
			indexSlice = append(indexSlice, i)
			validSIndexForChars[byteToIndex(s[i])] = indexSlice

			validSIndexes = append(validSIndexes, i) // ordered
		}
	}

	// if t has char not in s, or t has more duplicate chars
	isValid := true
	for i := range formatT {
		if formatT[i] > 0 && len(validSIndexForChars[i]) < formatT[i] {
			isValid = false
			break
		}
	}

	if !isValid {
		return ""
	}

	if len(t) == 1 {
		return t
	}

	firstValidWindow := make([]int, 0)
	for i := range validSIndexForChars {
		indexSlice := validSIndexForChars[i]
		for j := 0; j < formatT[i]; j++ { // index slice is valid after verify above
			firstValidWindow = append(firstValidWindow, indexSlice[0])
			indexSlice = indexSlice[1:]
		}
	}

	sort.Ints(firstValidWindow)
	start, end := firstValidWindow[0], firstValidWindow[len(firstValidWindow)-1]
	for i := start; i < len(validSIndexes); i = nextItem(validSIndexes, i) {
		indexSlice := validSIndexForChars[byteToIndex(s[i])]
		if len(indexSlice) <= 0 {
			break
		}

		formatSubS := make([]int, 52)
		for j := i; 0 <= j && j < len(validSIndexes); j = nextItem(validSIndexes, j) {
			formatSubS[byteToIndex(s[j])]++

			if isContains(formatT, formatSubS) {
				if j-i < end-start {
					start = i
					end = j
				}

				break
			}
		}
	}

	return s[start : end+1]
}

// byteToIndex rule: 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
func byteToIndex(char byte) int {
	index := -1

	if 'a' <= char && char <= 'z' {
		index = int(char - 'a')
	} else if 'A' <= char && char <= 'Z' {
		index = int(char - 'A' + 26)
	}

	return index
}

func nextItem(slice []int, target int) int {
	res := -1

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == target {
			res = slice[i+1]
			break
		}
	}

	return res
}

// isContains each a[i] <= b[i], i: [0, len(a)]
func isContains(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := range a {
		if a[i] > b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
