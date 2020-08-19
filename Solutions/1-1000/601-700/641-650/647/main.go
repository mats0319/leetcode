package mario

func countSubstrings(s string) int {
    if len(s) <= 1 {
        return 1
    }

    res := len(s)
    for i := 0; i < len(s) - 1; i++ {
        for l, r := i-1, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
            res++
        }

        if s[i] != s[i+1] {
            continue
        }

        for l, r := i, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
            res++
        }
    }

    return res
}
