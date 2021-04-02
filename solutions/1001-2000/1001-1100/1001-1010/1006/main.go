package mario

// sign loop: "*", "/", "+", "-", N >= 1
func clumsy(N int) int {
    stack := make([]int, 0, N)
    stack = append(stack, N)
    for i := N-1; i > 0; i-- {
        switch (N-i-1)%4 {
        case 0: // "*"
            stack[len(stack)-1] *= i
        case 1: // "/"
            stack[len(stack)-1] /= i
        case 2, 3: // "+", "-"
            stack = append(stack, i)
        }
    }

    res := stack[0]
    for i := 1; i < len(stack); i++ {
        if i & 1 == 1 { // "+"
            res += stack[i]
        } else {
            res -= stack[i]
        }
    }

    return res
}
