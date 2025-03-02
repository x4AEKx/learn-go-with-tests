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

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "new value"

		dictionary.Add("test", "new value")
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "existing"

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new value")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})

}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, expected string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertsString(t, got, expected)

}

func assertsString(t testing.TB, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}
