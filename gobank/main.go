package main

import (
	"fmt"

	"study.co.jp/gobank/accounts"
	"study.co.jp/gobank/customers"
)

func main() {

	account := accounts.CheckingAccount{
		Account: accounts.Account{
			Holder: customers.Holder{
				Name:       "Gandalfo",
				Document:   "SSID",
				Occupation: "Gray Wizard",
			},
			BranchNumber:  1,
			AccountNumber: 123,
		},
	}

	fmt.Println(account)

	account.Deposit(1000)
	fmt.Println(account.Balance())

	PayBill(&account, 100)
	fmt.Println(account.Balance())

	savingsAccount := accounts.SavingsAccount{}
	fmt.Println(savingsAccount)

}

func PayBill(account validateAccount, billAmount float64) {
	account.Withdraw(billAmount)
}

type validateAccount interface {
	Withdraw(amount float64) string
}
