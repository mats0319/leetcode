package mario

// s.length == t.length
// 题目只要求s->t单向1:1的映射关系
func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte, len(s))
	mReverse := make(map[byte]byte, len(s))

	isValid := true
	for i := range s {
		v, ok := m[s[i]]

		if !ok {
            _, isExist := mReverse[t[i]]
            if isExist {
                isValid = false
                break
            }

            m[s[i]] = t[i]
            mReverse[t[i]] = s[i]
        } else if t[i] != v {
            isValid = false
            break
        }
	}

	return isValid
}
