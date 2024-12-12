package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("world")
	want := "Hello world"
	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}
}
