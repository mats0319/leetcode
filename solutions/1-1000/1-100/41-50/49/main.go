package mario

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string, len(strs))
	for _, s := range strs {
		times := [26]int{}
		for i := 0; i < len(s); i++ {
			times[int(s[i]-'a')]++
		}

		m[times] = append(m[times], s)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
