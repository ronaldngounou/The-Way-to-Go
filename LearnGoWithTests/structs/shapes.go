package main 

import "math"
type Rectange struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}
// Syntax of a method: func (receiverName receiverType) MethodName(args)
// When doing this, we get a reference to the data via the receiverName variable.
// It is a convention in Go to have the receiver variable be the first letter of the type.

func (r Rectange) Area() float64 {
	return 2 * (r.Width + r.Height)
}



func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectange) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectange) float64 {
	return rectangle.Width * rectangle.Height
}
