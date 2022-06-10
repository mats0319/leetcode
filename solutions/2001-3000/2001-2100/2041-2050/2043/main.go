package mario

type Bank struct {
	balance []int64
	n       int
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: append([]int64{0}, balance...), // deprecated first item, use account number as index directly
		n:       len(balance),
	}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if !b.isValidAccountNumber(account1) || !b.isValidAccountNumber(account2) || b.balance[account1] < money {
        return false
    }

    b.balance[account1] -= money
    b.balance[account2] += money

    return true
}

func (b *Bank) Deposit(account int, money int64) bool {
    if !b.isValidAccountNumber(account) {
        return false
    }

    b.balance[account] += money

    return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
    if !b.isValidAccountNumber(account) || b.balance[account] < money {
        return false
    }

    b.balance[account] -= money

    return true
}

func (b *Bank) isValidAccountNumber(num int) bool {
	return 1 <= num && num <= b.n
}
