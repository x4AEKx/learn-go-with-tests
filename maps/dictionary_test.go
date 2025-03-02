package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "testing value"}

	t.Run("known value", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "testing value"
		assertsString(t, got, expected)
	})

	t.Run("", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertsString(t, err.Error(), expected)
	})
}

func assertsString(t testing.TB, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}
