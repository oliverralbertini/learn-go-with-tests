package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		key := "test"
		got, _ := dictionary.Search(key)
		want := "this is a test"

		assertStrings(t, want, got, key)
	})

	t.Run("unknown word", func(t *testing.T) {
		key := "tist"
		_, err := dictionary.Search(key)
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, want, err, key)
	})
}

func assertStrings(t *testing.T, want, got, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given '%s'", got, want, given)
	}
}

func assertError(t *testing.T, want, got error, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given '%s'", got, want, given)
	}
}
