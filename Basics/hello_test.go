package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jeanne")
	want := "Hello Jeanne"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
