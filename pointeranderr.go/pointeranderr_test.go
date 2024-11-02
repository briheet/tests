package pointeranderrgo

import (
	"golearn/utils"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		utils.AssertCorrectMessage(t, got, want)
	}
}
