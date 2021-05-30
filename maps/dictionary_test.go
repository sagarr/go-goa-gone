package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected error")
		}

		assertString(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add("test", "just a test")

		assertError(t, err, nil)
		assertDefination(t, dictionary, "test", "just a test")
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just a test"}

		err := dictionary.Add("test", "just a test")

		assertError(t, err, ErrWordExists)
		assertDefination(t, dictionary, "test", "just a test")
	})

}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	dictionary.Update("test", "new defination")

	assertDefination(t, dictionary, "test", "new defination")
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	dictionary.Delete("test")

	_, err := dictionary.Search("test")

	assertError(t, err, ErrNotFound)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertDefination(t testing.TB, dict Dictionary, word, defination string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find word", defination)
	}
	assertString(t, got, defination)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
