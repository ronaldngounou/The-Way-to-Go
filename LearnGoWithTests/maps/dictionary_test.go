package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertErrror(t, got, ErrNotFound)
	})

	
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("Marie", "1234")

	want := "0000"
	got, err := dictionary.Search("Marie")

	if err != nil {
		t.Fatal("Unable to find the word. Error happened: ", err)
	}

	assertStrings(t, got, want)

}





func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertErrror(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}