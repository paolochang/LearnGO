package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner string
	balance int
}

var errNotEnoughMoney = errors.New("not enough money to withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on the account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from the account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNotEnoughMoney
	}
	a.balance -= amount
	return nil
}

// Change owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Balance of the account
func (a Account) Balance() int {
	return a.balance
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Print format
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), ": ", a.Balance())
}