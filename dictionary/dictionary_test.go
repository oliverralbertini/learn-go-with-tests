package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test"}

	key := "test"
	got := Search(dictionary, key)
	want := "this is a test"

	assertStrings(t, want, got, key)
}

func assertStrings(t *testing.T, want, got, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given '%s'", got, want, given)
	}
}
