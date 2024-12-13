package Integer //this can be anyname we select, but other files in this folder
//gotta have the same name two, can't have two packages in a folder.
//Packages are used to unite different go files

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Basic addition", func(T *testing.T) {
		got := Add(2, 2)
		want := 4
		if got != want {
			t.Errorf("Wrong %d , got %d", want, got)
		}
	})
}
