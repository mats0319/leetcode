package mario

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for i := 1; i < len(strs) && len(res) > 0; i++ {
		if len(res) > len(strs[i]) {
			res = res[:len(strs[i])]
		}

		for j := range res {
			if strs[i][j] != res[j] {
				res = res[:j]
				break
			}
		}
	}

	return res
}
