package crypto

// give integer a specific name
type ZCash int

type Wallet struct {
	balance ZCash
}

// Just use * here to pass address, no need to dereference anywhere else
func (w *Wallet) Deposit(amount ZCash) {
	w.balance += amount
}

func (w *Wallet) Balance() ZCash {
	return w.balance
}
