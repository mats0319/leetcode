package mario

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := make([]int, 26)
    for i := range s {
        count[int(s[i]-'a')]++
        count[int(t[i]-'a')]--
    }

    i := 0
    for i < len(count) && count[i] == 0 {
        i++
    }

    return i == len(count)
}
