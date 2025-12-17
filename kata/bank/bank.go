package bank

import (
	"fmt"
	"io"
	"time"
)

type Account struct {
	actions []AccountAction
	writer  io.Writer
}

type AccountAction struct {
	date   time.Time
	amount int
}

func NewAccount(w io.Writer) *Account {
	return &Account{
		actions: []AccountAction{},
		writer:  w,
	}
}

func (a *Account) Deposit(amount int) {
	a.actions = append(a.actions, AccountAction{
		date:   time.Now().UTC(),
		amount: amount,
	})
}

func (a *Account) Withdraw(amount int) {
	a.actions = append(a.actions, AccountAction{
		date:   time.Now().UTC(),
		amount: amount * -1,
	})
}

func (a *Account) PrintStatement() {
	fmt.Fprintln(a.writer, "DATE | AMOUNT | BALANCE")
	balance := 0
	for _, action := range a.actions {
		balance += action.amount
		fmt.Fprintf(a.writer, "%s | %d | %d\n", action.date.String(), action.amount, balance)
	}
}
