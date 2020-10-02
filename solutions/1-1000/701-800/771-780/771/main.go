package mario

func numJewelsInStones(J string, S string) int {
    m := make(map[rune]bool, len(J))

    for _, ch := range J {
        m[ch] = true
    }

    count := 0
    for _, ch := range S {
        if m[ch] {
            count++
        }
    }

    return count
}
