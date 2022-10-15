package mario

func getPermutation(n int, k int) string {
    // n <= 9
    // list '1' to 'n', skip 'k-1' permutation, calc: k-1 = a*(n-1)!+b*(n-2)!+...+c*1!
    // a means first char need skip how many
    // e.g. [1,2,3,4] 24
    // 24-1 = 23 = 3*6+2*2+1*1
    // 1. first item skip 3: 4, list change to [1,2,3]
    // 2. second item skip 2: 3, list change to [1,2]
    // 3. third item skip 1: 2, list change to [1]
    // 4. res: [4,3,2,1]
    value := []int{-1, 1} // value: 'index'!, -1 for unity calc
    for i := 2; i < n; i++ {
        value = append(value, i*value[i-1])
    }

    resBytes := make([]byte, 0, n)
    used := make([]bool, n+1)
    k -= 1
    for i := 1; i <= n; i++ {
        skip := k / value[n-i]
        k -= skip * value[n-i]

        // find char and mark it
        for j := 1; j < len(used); j++ {
            if used[j] {
                continue
            }

            skip--
            if skip >= 0 {
                continue
            }

            used[j] = true
            resBytes = append(resBytes, byte(j)+'0')
            break
        }
    }

    return string(resBytes)
}
