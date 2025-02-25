package structs

type Rectangle struct {
	heigth float64
	width  float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.heigth + rectangle.width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.heigth * rectangle.width
}
