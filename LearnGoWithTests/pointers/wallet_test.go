package main 

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
	
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}

		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(20))

	})


	
}