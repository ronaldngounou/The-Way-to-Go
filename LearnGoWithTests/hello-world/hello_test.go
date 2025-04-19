package main 

import "testing"

/*

have (string)
want ()
 
We have to change the function Hello() to accept one argument.


*/
func TestHello(t *testing.T) {
	got := Hello("Christ")
	want := "Hello Christ"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
