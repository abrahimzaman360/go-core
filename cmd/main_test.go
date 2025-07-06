package main

import test "testing"

func TestHello(t *test.T) {
	t.Run("Hello", func(t *test.T) {
		_, err := sayHello("");

		if err != nil {
			t.Errorf("Error: Test unsuccessful!")
		}
	})
}