package mario

func uniquePaths(m int, n int) int {
    if m == 1 || n == 1 {
        return 1
    }

    return calculateCmn(m+n-2, min(m, n)-1)
}

func min(a, b int) (result int) {
    if a < b {
        result = a
    } else {
        result = b
    }

    return
}

func calculateCmn(m, n int) int {
    var mMR, nMR = 1, 1 // multiplication result
    for i := 1; i <= n; i, m = i+1, m-1 {
        mMR *= m
        nMR *= i
    }

    return mMR/nMR
}
