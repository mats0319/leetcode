package mario

func kthGrammar(N int, K int) int {
    if N == 1 {
        return 0
    }

    return kthGrammar(N-1, (K+1)/2) ^ ((K+1) % 2)
}
