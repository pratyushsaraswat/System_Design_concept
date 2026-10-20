package main

import "fmt"

type BankAccount struct {
	accountHolder string
	balance       float64
}

func NewBankAccount(accountHolder string, balance float64) *BankAccount {
	return &BankAccount{
		accountHolder: accountHolder,
		balance:       balance,
	}
}

func (a *BankAccount) Deposit(amount float64) {
	a.balance += amount
}

func (a *BankAccount) DisplayBalance() {
	fmt.Println(a.accountHolder, ":", a.balance)
}

func main() {
	a1 := NewBankAccount("Rahul", 1000)
	a2 := NewBankAccount("Priya", 5000)

	a1.Deposit(500)

	a1.DisplayBalance()
	a2.DisplayBalance()
}
