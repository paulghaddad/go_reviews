package pointers1

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) WithDraw(amount Bitcoin) {
	w.balance -= amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
