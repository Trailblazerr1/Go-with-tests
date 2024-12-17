package geometry

import "testing"

func TestArea(t *testing.T) {
	verifyResult := func(got, want float64, t *testing.T) {
		if got != want {
			t.Errorf("Got %.2f, want %.2f ", got, want)
		}
	}

	t.Run("Test area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 20}
		got := rectangle.Area()
		want := 200.0
		verifyResult(got, want, t)
	})
	//Go doesn't allows overloading of methods, so instead we have
	//to create methods, that is call Area on different instances
	//This is different from function
	t.Run("Test area of circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.15
		verifyResult(got, want, t)
	})
}
