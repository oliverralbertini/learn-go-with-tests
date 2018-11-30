package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' want '%s'\n", got, want)
		}
	}

	t.Run("saying hi to person", func(t *testing.T) {
		got := Hello("sb", "")
		want := "Hello, sb"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hi to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Elodie", "Italian")
		want := "Ciao, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
