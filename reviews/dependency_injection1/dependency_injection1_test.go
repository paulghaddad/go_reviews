package dependency_injection1

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Paul")

	got := buffer.String()
	want := "Hello, Paul"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
