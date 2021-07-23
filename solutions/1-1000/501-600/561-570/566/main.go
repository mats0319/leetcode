package mario

// 1 <= mat.length, mat[0].length <= 100
func matrixReshape(mat [][]int, r int, c int) [][]int {
    mr := len(mat)
    mc := len(mat[0])

    if mr * mc != r * c {
        return mat
    }

    data := make([]int, 0, r*c)
    for i := range mat {
        for j := range mat[i] {
            data = append(data, mat[i][j])
        }
    }

    index := 0
    res := make([][]int, r)
    for i := 0; i < len(res); i++ {
        res[i] = make([]int, c)

        for j := 0; j < len(res[i]); j++ {
            res[i][j] = data[index]
            index++
        }
    }

    return res
}
