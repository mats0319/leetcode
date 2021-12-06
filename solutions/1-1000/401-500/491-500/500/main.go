package mario

var lines = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

// findWords words not null, words[i] not null
func findWords(words []string) []string {
    res := make([]string, 0, len(words))
    for i := range words {
        index := getIndex(words[i][0])
        for j := 1; j < len(words[i]); j++ {
            if getIndex(words[i][j]) != index {
                index = -1
                break
            }
        }

        if index >= 0 {
            res = append(res, words[i])
        }
    }

    return res
}

func getIndex(char byte) int {
    if 'A' <= char && char <= 'Z' {
        char += 'a' - 'A'
    }

    index := -1
ALL:
    for i := 0; i < 3; i++ {
        for j := 0; j <len(lines[i]); j++ {
            if char == lines[i][j] {
                index = i
                break ALL
            }
        }
    }

    return index
}
