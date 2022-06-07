package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tony Stark", "")
		want := "Hello, Tony Stark"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Anthony Gonzalez", "Spanish")
		want := "Hola, Anthony Gonzalez"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Tanya T", "French")
		want := "Bonjour, Tanya T"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Bruno mars", "Portuguese")
		want := "Olá, Bruno mars"
		assertCorrectMessage(t, got, want)
	})
}
