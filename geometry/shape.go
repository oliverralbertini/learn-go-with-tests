package geometry

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Perimeter assumes an isosceles triangle
func (t *Triangle) Perimeter() float64 {
	return t.Base + 2.0*math.Sqrt(t.Height*t.Height+t.Base*t.Base/4.0)
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
