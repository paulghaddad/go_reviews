package structs_methods_interfaces1

import "math"

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	area() float64
}

func (rect Rectangle) area() float64 {
	return rect.length * rect.width
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (triangle Triangle) area() float64 {
	return 0.5 * triangle.base * triangle.height
}
