package structs_methods_interfaces1

import "testing"

// func TestPerimeter(t *testing.T) {
// 	t.Run("Rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{length: 10.0, width: 10.0}
//
// 		got := rectangle.perimeter()
// 		want := 40.0
//
// 		if got != want {
// 			t.Errorf("got: %f, want: %f", got, want)
// 		}
// 	})
//
// 	t.Run("Circle", func(t *testing.T) {
// 		circle := Circle{radius: 10.0}
//
// 		got := circle.perimeter()
// 		want := 62.83185307179586
//
// 		if got != want {
// 			t.Errorf("got: %g, want: %g", got, want)
// 		}
// 	})
// }
//
func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{length: 10.0, width: 10.0}, want: 100.0},
		{name: "Circle", shape: Circle{radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 10.0, height: 10.0}, want: 50.0},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.area()

		if got != want {
			t.Errorf("got: %f, want: %f", got, want)
		}
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}
}
