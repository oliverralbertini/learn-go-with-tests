package geometry

import (
	"math"
	"testing"
)

func compareFloats(t *testing.T, shape Shape, got, want float64) {
	t.Helper()

	if want != got {
		t.Errorf("#%v wanted '%.2f' got '%.2f'", shape, want, got)
	}
}

type ShapeTests []struct {
	shape   Shape
	want    float64
	subject string
}

func TestPerimeter(t *testing.T) {
	perimTests := ShapeTests{
		{shape: &Rectangle{Width: 1.0, Height: 2.0}, want: 6.0, subject: "perimeter of rectangle"},
		{shape: &Circle{Radius: 0.5}, want: math.Pi, subject: "perimeter of circle"},
		{shape: &Triangle{Base: 1.0, Height: math.Sqrt(3) / 2.0}, want: 3.0, subject: "perimeter of triangle"},
	}

	for _, pt := range perimTests {
		t.Run(pt.subject, func(t *testing.T) {
			compareFloats(t, pt.shape, pt.shape.Perimeter(), pt.want)
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := ShapeTests{
		{shape: &Rectangle{Width: 1.0, Height: 2.0}, want: 2.0, subject: "area of rectangle"},
		{shape: &Circle{Radius: 1.0}, want: math.Pi, subject: "area of circle"},
		{shape: &Triangle{Base: 1.0, Height: 2.0}, want: 1.0, subject: "area of triangle"},
	}

	for _, at := range areaTests {
		t.Run(at.subject, func(t *testing.T) {
			compareFloats(t, at.shape, at.shape.Area(), at.want)
		})
	}
}
