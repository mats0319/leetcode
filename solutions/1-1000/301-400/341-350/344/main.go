package mario

func reverseString(s []byte)  {
    if len(s) < 1 { // less than 2 is ok
        return
    }

    l, r := 0, len(s)-1
    for l < r {
        s[l], s[r] = s[r], s[l]

        l++
        r--
    }

    return
}
