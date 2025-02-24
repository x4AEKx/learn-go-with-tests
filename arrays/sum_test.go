package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("expected %d, but got %d, array: %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d, but got %d", expected, got)
	}
}
