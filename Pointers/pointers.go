package crypto

import (
	"errors"
)

// give integer a specific name
type ZCash int

type Wallet struct {
	balance ZCash
}

// var keyword allows us to have global variable
var ErrInsufficientFunds = errors.New("Insufficient funds")

// Just use * here to pass address, no need to dereference anywhere else
func (w *Wallet) Deposit(amount ZCash) {
	w.balance += amount
}

func (w *Wallet) Balance() ZCash {
	return w.balance
}

func (w *Wallet) Withdraw(amount ZCash) error {
	bal := w.Balance()
	if bal < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
