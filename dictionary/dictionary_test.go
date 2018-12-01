package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		key := "test"
		want := "this is a test"

		assertDefinition(t, dictionary, want, key)
	})

	t.Run("unknown word", func(t *testing.T) {
		key := "tist"
		_, err := dictionary.Search(key)
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want, key)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	key := "test"
	want := "this is just a test"
	dictionary.Add(key, want)

	assertDefinition(t, dictionary, want, key)
}

func assertDefinition(t *testing.T, dictionary Dictionary, want, word string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got '%s' want '%s' given the word '%s'", got, want, word)
	}
}

func assertError(t *testing.T, got, want error, word string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given the word '%s'", got, want, word)
	}
}
