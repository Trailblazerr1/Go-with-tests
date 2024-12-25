package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	verifyOutput := func(want, got string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	}

	verifyError := func(want, got error, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	}

	t.Run("Search in dict", func(t *testing.T) {
		dict := Dict{"key": "value"}
		got, _ := dict.Search("key")
		want := "value"
		verifyOutput(want, got, t)
	})

	t.Run("Search invalid value", func(t *testing.T) {
		dict := Dict{"key": "value"}
		_, err := dict.Search("key1")
		want := errorKeyNotFound

		if err == nil {
			t.Fatal("Should've thrown an error")
		}
		verifyError(want, err, t)
	})
}
