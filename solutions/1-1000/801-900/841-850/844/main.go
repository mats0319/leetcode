package mario

// str only contains low-case chars and '#'
func backspaceCompare(S string, T string) bool {
    isEqual := true
    for si, ti := len(S)-1, len(T)-1; si >= 0 || ti >= 0; {
        var sc, tc byte
        sc, si = getNextChar(S, si)
        tc, ti = getNextChar(T, ti)

        if sc != tc {
            isEqual = false
            break
        }
    }

    return isEqual
}

func getNextChar(str string, index int) (c byte, i int) {
    for count := 0; index >= 0; index-- {
        if str[index] == '#' {
            count++
        } else if count > 0 {
            count--
        } else {
            break // curr char is not '#', and no '#' left
        }
    }

    if index >= 0 {
        c = str[index]
        i = index-1
    } else if index < 0 {
        c = ' ' // str not contains space char
        i = -1
    }

    return
}
