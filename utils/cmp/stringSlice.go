package cmp

// CompareOnStringSlice compare if two string slice is equal
func CompareOnStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	var bMap = make(map[string]int, len(b))
	for i := range b {
		bMap[b[i]] = 0
	}

	for i := range a {
		if _, ok := bMap[a[i]]; ok {
			delete(bMap, a[i])
		} else {
			return false
		}
	}

	return len(bMap) == 0
}
