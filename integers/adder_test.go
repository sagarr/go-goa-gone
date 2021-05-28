package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(1, 2)
	expected := 3

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}