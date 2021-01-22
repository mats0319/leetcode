package mario

func findTheDifference(s string, t string) byte {
    times := make([]int, 26)
    for i := 0; i < len(s); i++ {
        times[s[i]-'a']--
        times[t[i]-'a']++
    }
    times[t[len(t)-1]-'a']++

    i := 0
    for i < len(times) && times[i] <= 0 {
        i++
    }

    return byte(i)+'a'
}
