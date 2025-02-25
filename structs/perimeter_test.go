package structs

import "testing"

func TestPerimetr(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f, but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	expected := 72.0

	if got != expected {
		t.Errorf("expected %.2f, but got %.2f", expected, got)
	}
}
