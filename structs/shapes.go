package structs

import "math"


type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Widhth float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Widhth)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Widhth
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Height float64
	Base float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * .5
}