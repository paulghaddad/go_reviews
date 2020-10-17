package structs_methods_interfaces

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

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) area() float64 {
	return .5 * t.base * t.height
}
