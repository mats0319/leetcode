package mario

var m map[int]int

func minDays(n int) int {
    if n == 0 || n == 1 {
        return n
    }

    var min int
    if v, ok := m[n-1]; ok {
        min = v
    } else {
        m[n-1] = minDays(n-1)
        min = m[n-1]
    }

    if n % 2 == 0 {
        v, ok := m[n/2]
        if !ok {
            v = minDays(n/2)
        }

        if min > v {
            min = v
        }
    }
    if n % 3 == 0 {
        v, ok := m[n/3]
        if !ok {
            v = minDays(n/3)
        }

        if min > v {
            min = v
        }
    }

    return min
}
