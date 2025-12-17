package bank

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountDeposit(t *testing.T) {
	buffer := &bytes.Buffer{}
	account := NewAccount(buffer)
	account.Deposit(10)

	assert.Len(t, account.actions, 1)
	assert.Equal(t, 10, account.actions[0].amount)
}

func TestAccountWithdraw(t *testing.T) {
	buffer := &bytes.Buffer{}
	account := NewAccount(buffer)
	account.Withdraw(10)

	assert.Len(t, account.actions, 1)
	assert.Equal(t, -10, account.actions[0].amount)
}

func TestAccountPrintStatement(t *testing.T) {
	buffer := &bytes.Buffer{}
	account := NewAccount(buffer)
	account.Deposit(150)
	account.Withdraw(20)

	account.PrintStatement()
	res := buffer.String()

	assert.Contains(t, res, "DATE | AMOUNT | BALANCE")
	assert.Contains(t, res, "150 | 150")
	assert.Contains(t, res, "-20 | 130")
}
