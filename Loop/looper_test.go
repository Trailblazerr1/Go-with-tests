package main

import "testing"

func TestLooper(t *testing.T) {
	t.Run("Base case , where something is repeated n no. of times", func(t *testing.T) {
		got := Looper("Hey ", 5)
		want := "Hey Hey Hey Hey Hey "
		if got != want {
			t.Errorf("Wanted %q got %q", want, got)
		}
	})
}
