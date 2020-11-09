package errors2

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got: %s; want: %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		fmt.Printf("address of balance in test is %v\n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		got := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))

		if got != nil {
			t.Errorf("Got an unexpected error")
		}
	})

	t.Run("Overdrawn funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		got := wallet.Withdraw(Bitcoin(50))
		want := OverdrawnFundsErr

		if got != want {
			t.Fatal("Expected an error")
		}

		expectedErrMessage := "Funds overdrawn"

		if got.Error() != expectedErrMessage {
			t.Errorf("got: %s; want: %s", got.Error(), expectedErrMessage)
		}
	})
}
