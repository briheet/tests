package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Perimeter() float64 {
	ans := 2 * (r.Height + r.Width)

	return ans
}

func (r *Rectangle) Area() float64 {
	ans := r.Height * r.Width
	return ans
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius * c.Radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
