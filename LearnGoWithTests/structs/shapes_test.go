package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectange{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestArea(t *testing.T) {
	t.Run("rectangles", func (t *testing.T) {
		rectangle := Rectange{10.0, 10.0}
		got := rectangle.Area() // we can call methods on "things"
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}		
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
} 
