package geometry

import "testing"

func TestArea(t *testing.T) {
	verifyResult := func(shape Shape, want float64, t *testing.T) {
		got := shape.Area()
		if got != want {
			t.Errorf("Got %.2f, want %.2f ", got, want)
		}
	}

	t.Run("Test area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 20}
		want := 200.0
		verifyResult(rectangle, want, t)
	})
	//Go doesn't allows overloading of methods, so instead we have
	//to create methods, that is call Area on different instances
	//This is different from function
	t.Run("Test area of circle", func(t *testing.T) {
		circle := Circle{10}
		want := 314.15
		verifyResult(circle, want, t)
	})
}

//Writing a table based tests for decoupling for shapes and tests

// func TestArea(t *testing T) {
// 	testParams := []struct {
// 		Shape
// 	}
// }
