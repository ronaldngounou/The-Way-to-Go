package main

import (
	"reflect"
	"testing"
)


func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}

		got := Sum(numbers)
		want := 150

		if got != want {
			t.Errorf("got %d, but want %d", got, want)
		}
	})
	
}


func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 4})
	want := []int{3, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, but want %d", got, want)
	}
}