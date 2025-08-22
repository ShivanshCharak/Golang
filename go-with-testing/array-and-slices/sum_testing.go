package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collected of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %v want %v given, %v", got, want, numbers)
		}

	})
	t.Run("collection of any size", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5}
		got := Sum(number)
		want := 6
		if got != want {
			t.Errorf("got %d want, %d given, %v", got, want, number)
		}
	})
}
