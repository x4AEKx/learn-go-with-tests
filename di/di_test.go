package di

import (
	"bytes"
	"testing"
)

func TestDi(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	expected := "Hello, Chris"

	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}
