package mario

import "fmt"

func getHint(secret string, guess string) string {
    resA, resB := 0, 0

    matched := make([]bool, len(secret))

    secMap := make(map[byte]int, len(secret))
    for i := range secret {
        if secret[i] != guess[i] {
            secMap[secret[i]]++
            continue
        }

        resA++
        matched[i] = true
    }

    for i := range guess {
        if matched[i] {
            continue
        }

        v, ok := secMap[guess[i]]
        if ok && v > 0 {
            secMap[guess[i]]--
            resB++
        }
    }

    return fmt.Sprintf("%dA%dB", resA, resB)
}
