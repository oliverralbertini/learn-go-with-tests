package geometry

import (
	"math"
	"testing"
)

func compareFloats(t *testing.T, got, want float64) {
	t.Helper()

	if want != got {
		t.Errorf("wanted '%.2f' got '%.2f'", want, got)
	}
}

func TestPerimeter(t *testing.T) {
	rect := Rectangle{1.0, 2.0}
	got := Perimeter(rect)
	want := 6.0

	compareFloats(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		compareFloats(t, got, want)
	}

	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{1.0, 2.0}

		checkArea(t, &rect, 2.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{1.0}

		checkArea(t, &circ, math.Pi)
	})
}
