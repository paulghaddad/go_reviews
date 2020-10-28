package errors1

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

		if got != nil {
			t.Fatalf("Error expected")
		}

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Overdrawn", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		got := wallet.Withdraw(Bitcoin(40))
		want := OverdrawnError

		assertError(t, got, want)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("Error expected")
	}

	if got != want {
		t.Errorf("got: %v; want: %v", got, want)
	}
}
