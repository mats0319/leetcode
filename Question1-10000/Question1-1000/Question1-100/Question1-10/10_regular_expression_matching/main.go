package rem

func isMatch(s string, p string) bool {
    isMatched := true
    for {
        if p == "" {
            isMatched = s == ""
            break
        }
        isMatched = REMatch(s, p)
        break
    }
    return isMatched
}

func REMatch(s, p string) (isMatched bool) {
    for {
        se := s == ""
        if p == "" {
            isMatched = se
            break
        }
        for ip := range p {
            if ip < len(p)-1 && p[ip+1] == '*' {
                if !se && (s[ip] == p[ip] || p[ip] == '.') {
                    isMatched = REMatch(s[1:], p) || REMatch(s, p[2:])
                } else {
                    isMatched = REMatch(s, p[2:])
                }
                break
            }
            if !se && (s[ip] == p[ip] || p[ip] == '.') {
                isMatched = REMatch(s[1:], p[1:])
            }
            break
        }
        break
    }
    return
}

// dynamic programming:
//
// s: [a-z]*
// n, m
// f(L, M): isMatch for suffix s[n-L:] and p[m-M:]
// f(L, M) = true, L == 0 && M == 0
//      if L == 0
//        | f(L, M-2), M > 2 && p[m-M+1] == "*"
//        | false
//      if M == 0
//        | false
//      if M > 1 && p[m-M+1] == "*"
//        |    f(L-1, M) || f(L-1, M-2), p[m-M] == s[n-L] || p[m-M] == '.'
//        |    f(L, M-2)
//      else
//        |    f(L - 1, M - 1), p[m-M] == "." || s[n-L] == p[m-M]
//        |    false
//
// -------------------------------------------------------------------------
//
func isMatch2(s string, p string) bool {

   str, pat := []rune(s), []rune(p)

   m, n := len(str), len(pat)
   dp := make([][]bool, m+1)
   for i := 0; i <= m; i++ {
       dp[i] = make([]bool, n+1)
   }

   for L := 0; L <= m; L++ {
       for M := 0; M <= n; M++ {
           if L == 0 && M == 0 {
               dp[L][M] = true
           } else if L == 0 {
               if M > 1 && pat[n-M+1] == '*'{
                   dp[L][M] = dp[L][M-2]
               } else {
                   dp[L][M] = false
               }
           } else if M == 0 {
               dp[L][M] = false
           } else if M > 1 && pat[n-M+1] == '*' {
               if pat[n-M] == str[m-L] || pat[n-M] == '.' {
                   dp[L][M] = dp[L-1][M] || dp[L][M-2]
               } else {
                   dp[L][M] = dp[L][M-2]
               }
           } else if pat[n-M] == '.' || pat[n-M] == str[m-L] {
               dp[L][M] = dp[L-1][M-1]
           }
       }
   }
   return dp[m][n]
}
