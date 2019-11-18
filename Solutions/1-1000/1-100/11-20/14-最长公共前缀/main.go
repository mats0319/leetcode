package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(result) > len(strs[i]) {
			result = result[:len(strs[i])]
		}

		if result == "" {
			break
		}

		for j := range result {
			if strs[i][j] != result[j] {
				result = result[:j]
				break
			}
		}
	}

	return result
}
