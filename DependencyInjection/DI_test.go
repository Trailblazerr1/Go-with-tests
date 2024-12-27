package main

import (
	"bytes"
	"testing"
)

func TestDI(t *testing.T) {
	t.Run("Test print's dependency injection", func(t *testing.T) {
		buffer := bytes.Buffer{}
		//this method will write its output to buffer instead of stdout
		Greet("World", &buffer)
		got := buffer.String()
		want := "Hello World"

		if got != want {
			t.Errorf("Wanted %q, got %q", want, got)
		}
	})
}
