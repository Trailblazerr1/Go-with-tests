package crypto

import "testing"

func TestWallet(t *testing.T) {
	verifyCases := func(wallet Wallet, want ZCash, t *testing.T) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	}

	verifyError := func(err error, t *testing.T, want error) {
		if err == nil {
			t.Fatal("Didn't return error while withdrawing insufficient funds")
		}
		//fatal breaks the flow and returns, so no null pointers
		if err != want {
			t.Errorf("Wanted %q got %q", want, err)
		}
	}

	verifyNoError := func(err error, t *testing.T) {
		if err != nil {
			t.Fatal("Didn't want any error, but got one")
		}
	}

	t.Run("Deposit money", func(t *testing.T) {
		wallet := Wallet{ZCash(0)}
		wallet.Deposit(ZCash(10))
		want := ZCash(10)
		verifyCases(wallet, want, t)
	})

	t.Run("Withdraw money", func(t *testing.T) {
		wallet := Wallet{ZCash(20)}
		err := wallet.Withdraw(ZCash(10))
		want := ZCash(10)
		verifyNoError(err, t)
		verifyCases(wallet, want, t)
	})

	t.Run("Withdraw money error", func(t *testing.T) {
		initialAmt := ZCash(20)
		wallet := Wallet{initialAmt}
		err := wallet.Withdraw(ZCash(50))
		verifyError(err, t, ErrInsufficientFunds)
		verifyCases(wallet, initialAmt, t) //verify if balance is samee after failure
		//of withdrawl
	})
}
