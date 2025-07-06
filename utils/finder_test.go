package utils

import (
	test "testing"
)

//TODO: Fix the tests if possible!
func TestFinder(t *test.T) {
	t.Run("Check Array Length", func(t *test.T) {
		arr := [5]int{1, 2, 3, 4, 5}

		res, err := Finder(arr[:])

		//TODO: Improvement is required!
		if res != len(arr) && err != nil {
			t.Error("Things are matching anyway!")
		}
	})
}