package mario

var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

// s.length >= 1
func modifyString(s string) string {
    res := []byte(s)
    for i := 0; i < len(res); i++ {
        if res[i] != '?' {
            continue
        }

        for j := 0; j < len(letters); j++ {
            if (i-1 < 0 || res[i-1] != letters[j]) && (i+1 >= len(res) || res[i+1] != letters[j]) {
                res[i] = letters[j]
                break
            }
        }
    }

    return string(res)
}
