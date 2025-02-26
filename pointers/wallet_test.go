package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	expected := 10

	if got != expected {
		t.Errorf("expected %d, but got %d, test is %p", expected, got, &wallet.balance)
	}

}
