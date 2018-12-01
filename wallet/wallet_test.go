package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet *Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	}

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("Was expecting an error, didn't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		want := Bitcoin(20)

		assertBalance(t, &wallet, want)
		assertError(t, err)
	})
}
