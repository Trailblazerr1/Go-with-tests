package dictionary

import (
	"fmt"
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

	verifyAdditionandUpdate := func(want, got bool, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("wanted true, got false")
		}
	}

	verifyDelete := func(want, err error, t *testing.T) {
		t.Helper()
		if err != want {
			t.Errorf("Delete failed")
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

	t.Run("Add key value in dict", func(t *testing.T) {
		dict := Dict{}
		added, err := dict.Add("key", "value")
		// want := true

		if err != nil {
			fmt.Print(err)
			t.Fatal("Shouldn't have thrown an error")
		}
		verifyAdditionandUpdate(added, true, t)
	})

	t.Run("Update value", func(t *testing.T) {
		dict := Dict{"Key": "Value"}
		err := dict.Update("Key", "newValue")
		if err != nil {
			fmt.Print(err)
			t.Fatal("Shouldn't have thrown an error")
		}
		// verifyAdditionandUpdate(added, true, t)
	})

	t.Run("Delete value", func(t *testing.T) {
		dict := Dict{"Key": "Value"}
		err := dict.Delete("Key")
		if err != nil {
			fmt.Print(err)
			t.Fatal("Shouldn't have thrown an error")
		}

		_, err = dict.Search("Key")
		want := errorKeyNotFound
		verifyDelete(want, err, t)
	})
}
