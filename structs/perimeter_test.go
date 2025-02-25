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
	areaTest := []struct {
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{12.0, 6.0}, expected: 72.0},
		{shape: Circle{10}, expected: 314.1592653589793},
		{shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, v := range areaTest {
		t.Run("", func(t *testing.T) {
			got := v.shape.Area()
			if got != v.expected {
				t.Errorf("%#v, expected %g, but got %g", v.shape, v.expected, got)
			}
		})
	}
}
