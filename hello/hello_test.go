package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Bonjour to people", func(t *testing.T) {
		got := Hello("Kimdir", "French")
		want := "Bonjour, Kimdir"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hola to people", func(t *testing.T) {
		got := Hello("Kimdir", "Spanish")
		want := "Hola, Kimdir"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hwllo, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q",got, want)
	}
}
