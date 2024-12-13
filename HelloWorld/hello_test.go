package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Generic test case", func(t *testing.T) {
		got := Hello("world", "English")
		want := "Hello world"
		checkCorrectness(got, want, t)
	})

	t.Run("When parameter is empty, it should default to 'Hello world' ", func(t *testing.T) {
		got := Hello("", "Hindi")
		want := "Namaste world"
		checkCorrectness(got, want, t)
	})
}

// testing.TB is an interface, hence it is used instead of testing.T as type
func checkCorrectness(got string, want string, t testing.TB) {
	t.Helper() //useful while checking error logs for better line numbers
	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}
}
