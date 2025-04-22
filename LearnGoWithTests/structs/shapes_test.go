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

// Table Driven Tests: https://go.dev/wiki/TableDrivenTests
func TestArea(t *testing.T) {
	tests := map[string] struct {
		input Shape
		result float64
	} {
		"rectangles": {
			input: Rectange{10.0, 10.0},
			result: 40.0,
		}, 
		"circles": {
			input: Circle{10},
			result: 314.1592653589793,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			if got, expected := test.input.Area(), test.result; got != expected {
				t.Fatalf("got %g, but expected %g", got, expected)
			}
		})
	}
}

