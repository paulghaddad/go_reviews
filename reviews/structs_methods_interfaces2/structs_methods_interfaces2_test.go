package structs_methods_interfaces2

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{length: 100, width: 100}

	got := rect.Perimeter()
	want := 400.0

	if got != want {
		t.Errorf("got: %g; want: %g", got, want)
	}
}

func TestArea(t *testing.T) {
	type test struct {
		name string
		shape Shape
		want float64
	}

	tests := []test{
		{name: "Rectangle", shape: Rectangle{length: 100, width: 100}, want: 10000.0},
		{name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
	}


	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got: %g; want: %g", got, want)
		}
	}

	for _, tc := range tests {
		checkArea(t, tc.shape, tc.want)
	}
}

