package mario

func checkPerfectNumber(num int) bool {
    sum := 1
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            sum += i
            if i*i < num {
                sum += num / i
            }
        }
    }

    return sum > 1 && sum == num
}
