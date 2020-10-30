package structs_methods_interfaces2

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length float64
	width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.length + 2*r.width
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c Circle) Perimeter() float64 {
	return 2*math.Pi*c.radius
}
