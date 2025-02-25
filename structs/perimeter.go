package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Heigth float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Heigth * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Heigth + rectangle.Width)
}
