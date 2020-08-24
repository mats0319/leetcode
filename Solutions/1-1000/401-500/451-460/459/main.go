package mario

func repeatedSubstringPattern(s string) bool {
    m := make(map[byte]int) // char-times

    for i := range s {
        m[s[i]]++
    }

    times := make([]int, 0, len(m))
    smallest := times[0]
    for _, v := range m {
        times = append(times, v)
        if v < smallest {
            smallest = v
        }
    }

    
}
