package main 

type Rectange struct {
	Width float64
	Height float64
}

func Perimeter(rectangle Rectange) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectange) float64 {
	return rectangle.Width * rectangle.Height
}
