package mario

import "strings"

// fullJustify every word.length <= maxWidth
func fullJustify(words []string, maxWidth int) []string {
    res := make([]string, 0)
    start := 0
    for start < len(words) {
        end := start + selectWords(words[start:], maxWidth)

        newLine := makeLine(words[start:end], maxWidth, end >= len(words))
        if len(newLine) < 1 {
            break
        }

        res = append(res, newLine)

        start = end
    }

    return res
}

func selectWords(words []string, maxWidth int) int {
    if len(words) < 1 {
        return 0
    }

    end := 0
    for length := -1; end < len(words); end++ {
        length += 1 + len(words[end]) // 1 is for the least blank between two words
        if length > maxWidth {
            break
        }
    }

    return end
}

func makeLine(words []string, maxWidth int, isLastLine bool) string {
    if len(words) <= 1 {
        for len(words[0]) < maxWidth {
            words[0] += " "
        }
        return words[0]
    }

    length := 0
    for i := range words {
        length += len(words[i])
    }

    res := ""
    if isLastLine {
        res = strings.Join(words, " ")
        for len(res) < maxWidth {
            res += " "
        }
    } else {
        for length < maxWidth {
            for i := 0; i < len(words)-1 && length < maxWidth; i++ {
                words[i] += " "
                length++
            }
        }

        res = strings.Join(words, "")
    }

    return res
}
