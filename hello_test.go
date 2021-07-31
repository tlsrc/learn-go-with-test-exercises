package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Tristan")
		want := "Hello Tristan"
		assertCorrectMessage(got, want, t)
	})

	t.Run("saying hello world when no input", func(t *testing.T) {
		got := hello("")
		want := "Hello world"
		assertCorrectMessage(got, want, t)
	})

}

func assertCorrectMessage(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
