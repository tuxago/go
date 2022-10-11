package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		b := []int{1, 2, 3}
		var got = sum(b)
		var want int = 6
		if got != want {
			t.Errorf("Expected %d, Got %d", want, got)
		}
	})
	t.Run("sum", func(t *testing.T) {
		b := []int{1, 2, 3, 4, 5}
		var got = sum(b)
		var want int = 15
		if got != want {
			t.Errorf("Expected %d, Got %d", want, got)
		}
	})
}
