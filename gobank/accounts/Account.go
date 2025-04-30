package accounts

import "study.co.jp/gobank/customers"

type Account struct {
	Holder                      customers.Holder
	BranchNumber, AccountNumber int
	balance                     float64 // Consider using a more robust type for currency
}

func (b *Account) Withdraw(amount float64) string {
	if amount > 0 && amount <= b.balance {
		b.balance -= amount
		return "Withdraw: Success"
	}
	return "Withdraw: Insufficient balance"
}

func (b *Account) Deposit(amount float64) string {
	if amount > 0 {
		b.balance += amount
		return "Deposit: Success"
	}
	return "Deposit: Invalid amount"
}

func (b *Account) Balance() float64 {
	return b.balance
}
