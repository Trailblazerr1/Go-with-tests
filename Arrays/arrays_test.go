package arrays

import (
	"reflect"
	"testing"
)

func TestArrays(t *testing.T) {
	t.Run("Return sum of array", func(t *testing.T) {
		//In go, size is part of array signature, so it must be same everywhere
		nos := [5]int{1, 2, 3, 4, 5}
		got := ArraySum(nos)
		want := 15

		if got != want {
			t.Errorf("Wanted %d, got %d for array %v", want, got, nos)
		}
	})

	t.Run("Sum of a slice", func(t *testing.T) {
		nos := []int{1, 2, 3, 4, 5}
		got := SliceSum(nos)
		want := 15

		if got != want {
			t.Errorf("Wanted %d, got %d for slice %v", want, got, nos)
		}
	})

	t.Run("Sum of multiple slice", func(t *testing.T) {
		got := SliceSumAll([]int{1, 2}, []int{3, 4, 5})
		want := []int{3, 12}
		//to compare array, use for loop or this
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted %d, got %d ", want, got)
		}
	})
}
