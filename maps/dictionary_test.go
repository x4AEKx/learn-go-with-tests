package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "testing value"}
	got := Search(dictionary, "test")
	expected := "testing value"

	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}
