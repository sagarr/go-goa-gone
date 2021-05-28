package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficent balance", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
	
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin)  {
	t.Helper()
	got := wallet.Balance()
	
	if want != got {
		t.Errorf("%s got, %s want", got, want)
	}
}

func assertError(t testing.TB, got, want error)  {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error")
	}

	if got != want {
		t.Errorf("%s want, %s got", want, got)
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Error("expected no error")
	}
}