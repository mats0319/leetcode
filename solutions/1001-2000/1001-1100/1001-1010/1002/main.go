package mario

func commonChars(A []string) []string {
	if len(A) <= 0 {
		return nil
	}

	min := [26]uint8{}
	curr := [26]uint8{}

	// first string
	for _, c := range A[0] {
		min[c-'a']++
	}

	for i := 1; i < len(A); i++ {
		for _, c := range A[i] {
			curr[c-'a']++
		}

		for i := range min {
			if curr[i] < min[i] {
				min[i] = curr[i]
			}

			curr[i] = 0
		}
	}

	res := make([]string, 0)
	var i, j uint8
	for ; i < uint8(len(min)); i++ {
		for j = 0; j < min[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}

	return res
}
