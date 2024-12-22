package crypto

import "testing"

func TestWallet(t *testing.T) {
	verifyCases := func(wallet Wallet, want ZCash, t *testing.T) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("Got %d, want %d", got, want)
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
		wallet.Withdraw(ZCash(10))
		want := ZCash(10)
		verifyCases(wallet, want, t)
	})
}
