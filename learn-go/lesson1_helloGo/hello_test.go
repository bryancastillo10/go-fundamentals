package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Gophers","Spanish")
		want := "Hola, Gophers"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("Bry","")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}