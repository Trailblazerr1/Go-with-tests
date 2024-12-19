package crypto

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit money", func(t *testing.T) {
		wallet := Wallet{ZCash(0)}
		wallet.Deposit(ZCash(10))
		got := wallet.Balance()
		want := ZCash(10)
		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}
