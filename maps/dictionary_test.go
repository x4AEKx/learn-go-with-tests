package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "testing value"}
	got := Search(dictionary, "test")
	expected := "testing value"
	assertsString(t, got, expected)

}

func assertsString(t testing.TB, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}
