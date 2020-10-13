package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := 15

	got := Sum(nums)

	if got != expected {
		t.Errorf("Sum(%v): %d; want: %d", nums, got, expected)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("SumAll(): %v; want: %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got []int, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got: %v; want: %v", got, expected)
		}
	}

	t.Run("Collections with elements", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 5})
		expected := []int{5, 14}

		checkSums(t, got, expected)
	})

	t.Run("Collections with no elements", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1}, []int{})
		expected := []int{5, 0, 0}

		checkSums(t, got, expected)
	})
}
