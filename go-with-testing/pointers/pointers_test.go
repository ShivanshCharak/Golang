package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	t.Run("Deposit", func(t *testing.T) {
		got := wallet.Deposit(400)
		want := "Deposited"
		if got != want {

			t.Errorf("deposit failed got %v want %v", got, want)
		}

		t.Logf("Depsot failed")

	})
	t.Run("Withdraw", func(t *testing.T) {
		got, message := wallet.Withdraw(20)
		if message != nil {
			t.Errorf(message.Error())
		}
		want := 380
		if got != want {
			t.Errorf("insuff funds got %d want %d", got, want)
		}

	})

}
