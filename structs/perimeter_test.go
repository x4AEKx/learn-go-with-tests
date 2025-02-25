package structs

import "testing"

func TestPerimetr(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f, but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	expected := 72.0

	if got != expected {
		t.Errorf("expected %.2f, but got %.2f", expected, got)
	}
}
