package fintech

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(wallet, Bitcoin(10), t)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(4))

		assertBalance(wallet, Bitcoin(6), t)
		assertNoError(err, t)
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		w := Wallet{balance: 5}

		err := w.Withdraw(20)

		assertBalance(w, 5, t)
		assertError(err, ErrInsufficientFunds, t)
	})

}

func assertNoError(got error, t *testing.T) {
	t.Helper()
	if got != nil {
		t.Errorf("Got an unwanted error %v", got)
	}
}

func assertError(got error, want error, t *testing.T) {
	t.Helper()
	if got == nil {
		t.Fatal("Did not receive an error")
	}

	if got != want {
		t.Errorf("Received an error with message '%s' instead of expected '%s'", got.Error(), want)
	}
}

func assertBalance(wallet Wallet, want Bitcoin, t *testing.T) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("Wanted balance of %s but got %s", want, got)
	}
}
