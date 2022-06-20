package mario

type ATM struct {
	quantity [5]int
	price    [5]int
}

func Constructor() ATM {
	return ATM{
		price: [5]int{20, 50, 100, 200, 500},
	}
}

func (a *ATM) Deposit(banknotesCount []int) {
	for i := range banknotesCount {
		a.quantity[i] += banknotesCount[i]
	}
}

func (a *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		q := small(a.quantity[i], amount/a.price[i])

		if q > 0 {
			amount -= q * a.price[i]
			res[i] = q
		}
	}

	if amount > 0 {
		return []int{-1}
	}

	for i := range a.quantity {
		a.quantity[i] -= res[i]
	}

	return res
}

func small(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
