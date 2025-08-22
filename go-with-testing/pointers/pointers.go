package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	amount        int
	accountNumber int
}

func (a *Wallet) Deposit(money int) string {
	a.amount += money
	fmt.Printf("Depsoted\n")
	return "Deposited"
}
func (a *Wallet) Balance() int {
	fmt.Println(a.amount)
	return a.amount
}

func (a *Wallet) Withdraw(amount int) (int, error) {
	if a.amount < amount {
		return 0, errors.New("insuff balance")
	}
	a.amount -= amount
	return a.amount, nil
}
