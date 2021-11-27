package mario

// s, p not null and only contains small-case letters
func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return nil
    }

    formatP := make([]int, 26)
    for i := range p {
        formatP[int(p[i]-'a')]++
    }

    res := make([]int, 0)
    formatSubS := make([]int, 26)
    reFormatSubS := true
    for start := 0; start <= len(s)-len(p); start++ {
        if reFormatSubS {
            reFormatSubS = false
            for traversal := start; traversal < start + len(p); traversal++ {
                if formatP[int(s[traversal]-'a')] == 0 {
                    formatSubS = make([]int, 26)
                    start = traversal // start will add one outside
                    reFormatSubS = true
                    break
                }

                formatSubS[int(s[traversal]-'a')]++
            }

            if reFormatSubS {
                continue
            }
        } else {
            formatSubS[int(s[start-1]-'a')]--
            formatSubS[int(s[start+len(p)-1]-'a')]++
        }

        if isStrictEqual(formatP, formatSubS) {
            res = append(res, start)
        }
    }

    return res
}

func isStrictEqual(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    isEqual := true
    for i := range a {
        if a[i] != b[i] {
            isEqual = false
            break
        }
    }

    return isEqual
}
