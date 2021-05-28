package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func (t testing.TB, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sagar", "english")
		want := "Hello, Sagar"
		assertMessage(t, got, want)
	})
	
	t.Run("say 'Hello, World' when empty name passed", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertMessage(t, got, want)
	})

}