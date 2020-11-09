package errors2

import (
	"errors"
	"fmt"
)

type Bitcoin int

var OverdrawnFundsErr = errors.New("Funds overdrawn")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return OverdrawnFundsErr
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Stringer interface {
	String() string
}
