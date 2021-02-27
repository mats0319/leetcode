package mario

// A.length >= K
func subarraysWithKDistinct(A []int, K int) int {
    count := 0
    for i := 0; i < len(A); i++ {
        m := make(map[int]int, len(A)-i)

        j := i
        for ; j < len(A) && len(m) < K; j++ {
            m[A[j]]++
        }

        if len(m) < K {
            break
        }

        count++
        for ; j < len(A); j++ {
            if _, ok := m[A[j]]; !ok {
                break
            }

            count++
        }
    }

    return count
}
