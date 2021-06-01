package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sagar")

	got := buffer.String()
	want := "Hello Sagar"

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}