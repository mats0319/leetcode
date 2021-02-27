package mario

func generateParenthesis(n int) []string {
    return dfs("", n, n, make([]string, 0))
}

func dfs(currStr string, left, right int, res []string) []string {
    if left == 0 && right == 0 {
        return append(res, currStr)
    }

    if left > right {
        return res
    }

    if left > 0 {
        res = dfs(currStr+"(", left-1, right, res)
    }
    if right > 0 {
        res = dfs(currStr+")", left, right-1, res)
    }

    return res
}
